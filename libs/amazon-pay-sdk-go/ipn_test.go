package amazonpay

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIPN(t *testing.T) {
	msg := `{
  "Type" : "Notification",
  "MessageId" : "f2733eaf-c284-5eda-bd2b-6d857d1aefd0",
  "TopicArn" : "arn:aws:sns:us-west-2:291180941288:A31YDYE76E6TCPA3G30QK8ZESA0D",
  "Message" : "{\"MarketplaceID\":\"A31YDYE76E6TCP\",\"ReleaseEnvironment\":\"Sandbox\",\"Version\":\"2013-01-01\",\"NotificationType\":\"OrderReferenceNotification\",\"SellerId\":\"A3G30QK8ZESA0D\",\"NotificationReferenceId\":\"1111111-1111-11111-1111-11111EXAMPLE\",\"IsSample\":true,\"Timestamp\":\"2018-03-20T07:50:54.554Z\",\"NotificationData\":\"<?xml version=\\\"1.0\\\" encoding=\\\"UTF-8\\\"?>\\n            <OrderReferenceNotification xmlns=\\\"https://mws.amazonservices.com/ipn/OffAmazonPayments/2013-01-01\\\">\\n                <OrderReference>\\n                    <AmazonOrderReferenceId>P01-0000000-0000000-000000<\\/AmazonOrderReferenceId>\\n                    <OrderTotal>\\n                        <Amount>0.0<\\/Amount>\\n <CurrencyCode>USD<\\/CurrencyCode>\\n                    <\\/OrderTotal>\\n                    <SellerOrderAttributes />\\n <OrderReferenceStatus>\\n                        <State>Closed<\\/State>           \\n                        <LastUpdateTimestamp>2013-01-01T01:01:01.001Z<\\/LastUpdateTimestamp>\\n                        <ReasonCode>AmazonClosed<\\/ReasonCode>\\n                    <\\/OrderReferenceStatus>\\n                    <CreationTimestamp>2013-01-01T01:01:01.001Z<\\/CreationTimestamp>       \\n                    <ExpirationTimestamp>2013-01-01T01:01:01.001Z<\\/ExpirationTimestamp>\\n                <\\/OrderReference>\\n            <\\/OrderReferenceNotification>\"}",
  "Timestamp" : "2018-03-20T07:50:54.642Z",
  "SignatureVersion" : "1",
  "Signature" : "Z1ETrmSHJfc5vPgYLi+wdXumCcQUajPjNrt2RwOp1LO4zHIYt+houIo/RI2smYCxiFyUzZBdDPixv1HOwcpm6l4LciC08JkORy/QrHWZmPpvQPgTo6wXBpXh61z2F5Qb7Pcn3z2BYrNyad7vVHPbLs0bqaXqEZwWr9EHq6BF94vwQZBzdTnp84ISb26rh91AgHYMWKVQaTopme6hfJSGToWUrFoMklF7Vt+GPkoPcjhQX6sGuBm9RC4U0VLrv1kj3D7qXjl2lPfkR7Fd5jft9Y2qsfRjObgHfOEg71vsELEjzhClo/KTcA1mp2YlRCzNBWdrA6tbIDpMzsXa1NGb6w==",
  "SigningCertURL" : "https://sns.us-west-2.amazonaws.com/SimpleNotificationService-433026a4050d206028891664da859041.pem",
  "UnsubscribeURL" : "https://sns.us-west-2.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-west-2:291180941288:A31YDYE76E6TCPA3G30QK8ZESA0D:5a3fd7d6-518b-4266-a749-00e086105270"
}`

	var ipn IPN
	if err := json.Unmarshal([]byte(msg), &ipn); err == nil {
		if msg, err := ipn.GetMessage(); err == nil {
			fmt.Println(msg)
			if data, err := msg.GetOrderReferenceNotification(); err == nil {
				fmt.Printf("%#v \n", data)
			} else {
				t.Errorf("No error should happen when get order reference notification, but got %v", err)
			}
		} else {
			t.Errorf("No error should happen when get message, but got %v", err)
		}
	} else {
		t.Errorf("No error should happen when unmarshal ipn message, but got %v", err)
	}
}

func TestPem(t *testing.T) {
	cert := getCert(&IPN{SigningCertURL: "https://sns.us-west-2.amazonaws.com/SimpleNotificationService-433026a4050d206028891664da859041.pem"})

	if cert == nil {
		t.Error("No Error should happen when download cert")
	} else if !verifyCertSubject(cert) {
		t.Error("Cert's subject should be correct")
	}
}
