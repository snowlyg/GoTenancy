package paygent

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"GoTenancy/libs/gomerchant"
)

type Paygent struct {
	Config *Config
}

type Config struct {
	MerchantID      string `required:"true"`
	ConnectID       string `required:"true"`
	ConnectPassword string `required:"true"`
	TelegramVersion string

	CertPassword      string
	ClientFilePath    string // this is required, if ClientFileContent is blank
	ClientFileContent string // this is required, if ClientFilePath is blank
	CAFilePath        string // this is required, if CAFileContent is blank
	CAFileContent     string // this is required, if CAFilePath is blank

	ProductionMode bool
}

func New(config *Config) *Paygent {
	return &Paygent{
		Config: config,
	}
}

func (paygent *Paygent) Client() (*http.Client, error) {
	var (
		certBytes, certKeyBytes []byte
		caCertPool              = x509.NewCertPool()
	)

	if paygent.Config.ClientFileContent == "" {
		if pemData, err := ioutil.ReadFile(paygent.Config.ClientFilePath); err == nil {
			paygent.Config.ClientFileContent = string(pemData)
		} else {
			return nil, err
		}
	}

	if paygent.Config.CAFileContent == "" {
		if pemData, err := ioutil.ReadFile(paygent.Config.CAFilePath); err == nil {
			paygent.Config.CAFileContent = string(pemData)
		} else {
			return nil, err
		}
	}

	{ // ClientFile
		var (
			originalPemData []byte
			block           *pem.Block
			pemData         = []byte(paygent.Config.ClientFileContent)
		)

		caCertPool.AppendCertsFromPEM(pemData)

		for len(pemData) > 0 {
			originalPemData = pemData
			if block, pemData = pem.Decode(pemData); block == nil {
				break
			}

			if block.Type == "CERTIFICATE" && len(certBytes) == 0 {
				if len(certBytes) == 0 {
					certBytes = originalPemData[0 : len(originalPemData)-len(pemData)-1]
				}
			}

			if block.Type == "RSA PRIVATE KEY" {
				if x509.IsEncryptedPEMBlock(block) {
					if results, err := x509.DecryptPEMBlock(block, []byte(paygent.Config.CertPassword)); err == nil {
						certKeyBytes = []byte("-----BEGIN RSA PRIVATE KEY-----\n" + base64.StdEncoding.EncodeToString(results) + "\n-----END RSA PRIVATE KEY-----")
					} else {
						return nil, err
					}
				} else {
					certKeyBytes = originalPemData[0 : len(originalPemData)-len(pemData)-1]
				}
			}
		}
	}

	{ // CAFile
		caCertPool.AppendCertsFromPEM([]byte(paygent.Config.CAFileContent))
	}

	if cert, err := tls.X509KeyPair(certBytes, certKeyBytes); err == nil {
		// Setup HTTPS client
		tlsConfig := &tls.Config{
			Certificates:  []tls.Certificate{cert},
			RootCAs:       caCertPool,
			Renegotiation: tls.RenegotiateFreelyAsClient,
		}
		tlsConfig.BuildNameToCertificate()
		transport := &http.Transport{TLSClientConfig: tlsConfig}
		cookieJar, _ := cookiejar.New(nil)

		return &http.Client{
			Transport: transport,
			Jar:       cookieJar,
		}, nil
	} else {
		return nil, err
	}
}

func (paygent *Paygent) serviceURLOfTelegramKind(telegramKind string) (*url.URL, error) {
	var (
		domain  = TelegramServiceSandboxDomain
		urlPath string
	)

	if paygent.Config.ProductionMode {
		domain = TelegramServiceDomain
	}

	for i := 0; i < len(telegramKind)-1; i++ {
		if p, ok := TelegramServiceURLs[telegramKind[0:len(telegramKind)-i]]; ok {
			urlPath = p
			break
		}
	}

	u, err := url.Parse(domain)
	u.Path = urlPath
	return u, err
}

var ResponseParser = regexp.MustCompile(`(?s)(\w+?)=(<!DOCTYPE.*HTML>|.*?)(\r\n|$)`)

type Response struct {
	RawBody        string
	Result         string
	ResponseCode   string
	ResponseDetail string
	gomerchant.Params
}

func (paygent *Paygent) Request(telegramKind string, params gomerchant.Params) (Response, error) {
	var (
		response    *http.Response
		serviceURL  *url.URL
		urlValues   = url.Values{}
		results     = Response{Params: gomerchant.Params{}}
		client, err = paygent.Client()
	)

	if err == nil {
		serviceURL, err = paygent.serviceURLOfTelegramKind(telegramKind)

		if err == nil {
			urlValues.Add("merchant_id", paygent.Config.MerchantID)
			urlValues.Add("connect_id", paygent.Config.ConnectID)
			urlValues.Add("connect_password", paygent.Config.ConnectPassword)
			if paygent.Config.TelegramVersion != "" {
				urlValues.Add("telegram_version", paygent.Config.TelegramVersion)
			} else {
				urlValues.Add("telegram_version", "1.0")
			}
			urlValues.Add("telegram_kind", telegramKind)

			for key, value := range params {
				urlValues.Add(key, fmt.Sprint(value))
			}

			response, err = client.Post(serviceURL.String(), "application/x-www-form-urlencoded", strings.NewReader(urlValues.Encode()))

			if err == nil {
				if response.StatusCode == 200 {
					defer response.Body.Close()
					var bodyBytes []byte
					bodyBytes, err = ioutil.ReadAll(response.Body)

					shiftJISToUTF8 := transform.NewReader(bytes.NewReader(bodyBytes), japanese.ShiftJIS.NewDecoder())
					utf8Bytes, _ := ioutil.ReadAll(shiftJISToUTF8)
					results.RawBody = string(utf8Bytes)
					results.Params.Set("RawBody", results.RawBody)

					if err == nil {
						for _, value := range ResponseParser.FindAllStringSubmatch(string(utf8Bytes), -1) {
							if value[1] == "result" {
								results.Result = value[2]
							}

							if value[1] == "response_code" {
								results.ResponseCode = value[2]
							}

							if value[1] == "response_detail" {
								results.ResponseDetail = value[2]
							}

							results.Set(value[1], value[2])
						}

						if results.Result == "1" {
							if results.ResponseDetail != "" {
								err = errors.New(results.ResponseDetail)
							} else {
								err = errors.New("failed to process this request")
							}
						}

						return results, err
					}
				}
				err = fmt.Errorf("status code: %v", response.StatusCode)
			}
		}
	}

	return results, err
}

func (paygent *Paygent) Authorize(amount uint64, params gomerchant.AuthorizeParams) (gomerchant.AuthorizeResponse, error) {
	var (
		response      gomerchant.AuthorizeResponse
		requestParams = gomerchant.Params{
			"trading_id":     params.OrderID,
			"payment_amount": amount,
			"payment_class":  10,
		}
	)

	if ok, threeDomainParams := get3DModeParams(params); ok {
		requestParams["http_user_agent"] = threeDomainParams.UserAgent
		requestParams["term_url"] = threeDomainParams.TermURL
		requestParams["http_accept"] = threeDomainParams.HttpAccept
	} else {
		requestParams["3dsecure_ryaku"] = 1
	}

	if paymentMethod := params.PaymentMethod; paymentMethod != nil {
		if savedCreditCard := paymentMethod.SavedCreditCard; savedCreditCard != nil {
			requestParams["stock_card_mode"] = 1
			requestParams["customer_id"] = savedCreditCard.CustomerID
			requestParams["customer_card_id"] = savedCreditCard.CreditCardID
		} else if creditCard := paymentMethod.CreditCard; creditCard != nil {
			requestParams["card_number"] = creditCard.Number
			requestParams["card_valid_term"] = getValidTerm(creditCard)
		} else {
			return response, gomerchant.ErrNotSupportedPaymentMethod
		}
	} else {
		return response, gomerchant.ErrNotSupportedPaymentMethod
	}

	results, err := paygent.Request("020", requestParams)
	if err == nil {
		if paymentID, ok := results.Get("payment_id"); ok {
			response.TransactionID = fmt.Sprint(paymentID)
		}

		// If 3D Mode
		if ok, _ := get3DModeParams(params); ok {
			if result, ok := results.Get("out_acs_html"); ok && fmt.Sprint(result) != "" {
				response.HandleRequest = true
				response.RequestHandler = func(writer http.ResponseWriter, request *http.Request, _ gomerchant.Params) error {
					_, e := io.WriteString(writer, fmt.Sprint(result))
					return e
				}
			}
		}
	}

	response.Params = results.Params

	return response, err
}

func (paygent *Paygent) Capture(transactionID string, params gomerchant.CaptureParams) (gomerchant.CaptureResponse, error) {
	var (
		response      gomerchant.CaptureResponse
		requestParams = gomerchant.Params{"payment_id": transactionID}
	)

	results, err := paygent.Request("022", requestParams)

	if err == nil {
		if paymentID, ok := results.Get("payment_id"); ok {
			response.TransactionID = fmt.Sprint(paymentID)
		}
	}
	response.Params = results.Params

	return response, err
}

func (paygent *Paygent) Refund(transactionID string, amount uint, params gomerchant.RefundParams) (response gomerchant.RefundResponse, err error) {
	var (
		results       Response
		requestParams = gomerchant.Params{
			"payment_id":     transactionID,
			"payment_amount": amount,
			"reduction_flag": 1,
		}
	)

	if params.Captured {
		results, err = paygent.Request("029", requestParams)
	} else {
		results, err = paygent.Request("028", requestParams)
	}

	response.Params = results.Params
	if paymentID, ok := getPaymentID(results); ok {
		response.TransactionID = paymentID
	}

	return response, err
}

func (paygent *Paygent) Void(transactionID string, params gomerchant.VoidParams) (response gomerchant.VoidResponse, err error) {
	var results Response

	if params.Captured {
		results, err = paygent.Request("023", gomerchant.Params{"payment_id": transactionID})
	} else {
		results, err = paygent.Request("021", gomerchant.Params{"payment_id": transactionID})
	}

	response.Params = results.Params
	if paymentID, ok := getPaymentID(results); ok {
		response.TransactionID = paymentID
	}

	return response, err
}

func (paygent *Paygent) Query(transactionID string) (gomerchant.Transaction, error) {
	results, err := paygent.Request("094", gomerchant.Params{"payment_id": transactionID})
	transaction := extractTransactionFromPaygentResponse(results)
	transaction.Params = results.Params
	return transaction, err
}
