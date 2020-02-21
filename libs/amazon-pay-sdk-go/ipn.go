package amazonpay

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// IPN ipn message
type IPN struct {
	Type             string
	MessageID        string `json:"MessageId"`
	TopicArn         string
	Timestamp        string
	SignatureVersion string
	Signature        string
	SigningCertURL   string
	UnsubscribeURL   string
	Message          string
}

// Message ipn message struct
type Message struct {
	NotificationReferenceID string `json:"NotificationReferenceId"`
	NotificationType        string
	SellerID                string `json:"SellerId"`
	ReleaseEnvironment      string
	Version                 string
	NotificationData        string
	Timestamp               string
}

// GetMessage get IPN message
func (ipn IPN) GetMessage() (Message, error) {
	var msg Message
	err := json.Unmarshal([]byte(ipn.Message), &msg)
	return msg, err
}

// GetNotification get notification message
func (ipn IPN) GetNotification() (interface{}, error) {
	msg, err := ipn.GetMessage()

	if err == nil {
		switch msg.NotificationType {
		case "OrderReferenceNotification":
			return msg.GetOrderReferenceNotification()
		case "PaymentAuthorize":
			return msg.GetAuthorizationNotification()
		case "PaymentCapture":
			return msg.GetCaptureNotification()
		case "PaymentRefund":
			return msg.GetRefundNotification()
		}
	}

	return nil, err
}

// GetOrderReferenceNotification get order reference notification data (notification type: OrderReferenceNotification)
func (msg Message) GetOrderReferenceNotification() (OrderReferenceNotification, error) {
	var notification OrderReferenceNotification
	err := xml.Unmarshal([]byte(msg.NotificationData), &notification)
	return notification, err
}

// GetAuthorizationNotification get authorization notification data (notification type: PaymentAuthorize)
func (msg Message) GetAuthorizationNotification() (AuthorizationNotification, error) {
	var notification AuthorizationNotification
	err := xml.Unmarshal([]byte(msg.NotificationData), &notification)
	return notification, err
}

// GetCaptureNotification get capture notification data (notification type: PaymentCapture)
func (msg Message) GetCaptureNotification() (CaptureNotification, error) {
	var notification CaptureNotification
	err := xml.Unmarshal([]byte(msg.NotificationData), &notification)
	return notification, err
}

// GetRefundNotification get refund notification data (notification type: PaymentRefund)
func (msg Message) GetRefundNotification() (RefundNotification, error) {
	var notification RefundNotification
	err := xml.Unmarshal([]byte(msg.NotificationData), &notification)
	return notification, err
}

// VerifyIPNRequest verify IPN request message
func VerifyIPNRequest(req *http.Request) (*IPN, bool) {
	if req.Header.Get("x-amz-sns-message-type") != "Notification" {
		return nil, false
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, false
	}

	var ipn IPN
	if err := json.Unmarshal(body, &ipn); err != nil {
		return nil, false
	}

	if cert := getCert(&ipn); cert != nil && verifyCertSubject(cert) && verifySignedString(&ipn, cert) {
		return &ipn, true
	}

	return &ipn, false
}

func getCert(ipn *IPN) *x509.Certificate {
	if resp, err := http.Get(ipn.SigningCertURL); err == nil {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			block, _ := pem.Decode([]byte(body))
			if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
				return cert
			}
		}
	}
	return nil
}

func verifyCertSubject(cert *x509.Certificate) bool {
	return cert.Subject.CommonName == "sns.amazonaws.com"
}

func verifySignedString(ipn *IPN, cert *x509.Certificate) bool {
	canonicalString := fmt.Sprintf("")

	if ipn.Message != "" {
		canonicalString += "Message\n" + ipn.Message + "\n"
	}

	if ipn.MessageID != "" {
		canonicalString += "MessageId\n" + ipn.MessageID + "\n"
	}

	if ipn.Timestamp != "" {
		canonicalString += "Timestamp\n" + ipn.Timestamp + "\n"
	}

	if ipn.TopicArn != "" {
		canonicalString += "TopicArn\n" + ipn.TopicArn + "\n"
	}

	if ipn.Type != "" {
		canonicalString += "Type\n" + ipn.Type + "\n"
	}

	ds, _ := base64.StdEncoding.DecodeString(ipn.Signature)
	h := sha1.New()
	h.Write([]byte(canonicalString))
	digest := h.Sum(nil)
	return rsa.VerifyPKCS1v15(cert.PublicKey.(*rsa.PublicKey), crypto.SHA1, digest, ds) == nil
}
