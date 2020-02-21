package paygent_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"GoTenancy/libs/gomerchant"
	"GoTenancy/libs/gomerchant/gateways/paygent"
	"GoTenancy/libs/gomerchant/tests"
	"github.com/jinzhu/configor"
)

var Paygent *paygent.Paygent

type Config struct {
	MerchantID      string `required:"true"`
	ConnectID       string `required:"true"`
	ConnectPassword string `required:"true"`
	TelegramVersion string `required:"true" default:"1.0"`

	ClientFilePath string `required:"true" default:"paygent.pem"`
	CertPassword   string `required:"true" default:"changeit"`
	CAFilePath     string `required:"true" default:"curl-ca-bundle.crt"`

	ProductionMode bool
}

func init() {
	var config = &Config{}
	os.Setenv("CONFIGOR_ENV_PREFIX", "-")
	if err := configor.Load(config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Paygent = paygent.New(&paygent.Config{
		MerchantID:      config.MerchantID,
		ConnectID:       config.ConnectID,
		ConnectPassword: config.ConnectPassword,
		ClientFilePath:  config.ClientFilePath,
		CertPassword:    config.CertPassword,
		CAFilePath:      config.CAFilePath,
		ProductionMode:  config.ProductionMode,
	})
}

func TestTestSuite(t *testing.T) {
	tests.TestSuite{
		CreditCardManager: Paygent,
		Gateway:           Paygent,
		GetRandomCustomerID: func() string {
			return fmt.Sprint(time.Now().Unix())
		},
	}.TestAll(t)
}

func Test3DAuthorizeAndCapture(t *testing.T) {
	cards := map[string]bool{
		"5123459358515820": true,
		"5123459358515821": false,
	}

	for card, is3D := range cards {
		authorizeResult, err := Paygent.SecureCodeAuthorize(100,
			paygent.SecureCodeParams{
				UserAgent: "User-Agent	Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12) AppleWebKit/602.3.12 (KHTML, like Gecko) Version/10.0.2 Safari/602.3.12",
				TermURL:    "http://getqor.com/order/return",
				HttpAccept: "http",
			},
			gomerchant.AuthorizeParams{
				Currency: "JPY",
				OrderID:  fmt.Sprint(time.Now().Unix()),
				PaymentMethod: &gomerchant.PaymentMethod{
					CreditCard: &gomerchant.CreditCard{
						Name:     "JCB Card",
						Number:   card,
						ExpMonth: 1,
						ExpYear:  uint(time.Now().Year() + 1),
					},
				},
			})

		if err != nil || authorizeResult.TransactionID == "" {
			t.Error(err, authorizeResult)
		}

		if is3D != authorizeResult.HandleRequest {
			t.Errorf("HandleRequest for card %v should be %v, but got %v", card, is3D, authorizeResult.HandleRequest)
		}

		if is3D {
			if result, ok := authorizeResult.Get("out_acs_html"); !ok || result.(string) == "" {
				t.Errorf("should get HTML, but %v", authorizeResult)
			}
		}
	}
}
