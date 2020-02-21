package amazonpay

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestOrderReferenceDetails(t *testing.T) {
	resp := `<GetOrderReferenceDetailsResponse xmlns="http://mws.amazonservices.com/schema/OffAmazonPayments/2013-01-01">
  <GetOrderReferenceDetailsResult>
    <OrderReferenceDetails>
      <OrderReferenceStatus>
        <State>Draft</State>
      </OrderReferenceStatus>
      <Destination>
        <DestinationType>Physical</DestinationType>
        <PhysicalDestination>
          <StateOrRegion>Tokyo</StateOrRegion>
          <City>Ota-ku</City>
          <CountryCode>JP</CountryCode>
          <PostalCode>144-8588</PostalCode>
        </PhysicalDestination>
      </Destination>
      <ExpirationTimestamp>2018-08-28T16:00:23.013Z</ExpirationTimestamp>
      <IdList/>
      <SellerOrderAttributes/>
      <OrderTotal>
        <CurrencyCode>JPY</CurrencyCode>
        <Amount>23.00</Amount>
      </OrderTotal>
      <ReleaseEnvironment>Sandbox</ReleaseEnvironment>
      <AmazonOrderReferenceId>S03-1594298-4362493</AmazonOrderReferenceId>
      <CreationTimestamp>2018-03-01T16:00:23.013Z</CreationTimestamp>
      <RequestPaymentAuthorization>false</RequestPaymentAuthorization>
    </OrderReferenceDetails>
  </GetOrderReferenceDetailsResult>
  <ResponseMetadata>
    <RequestId>ae47aef3-e030-4fdf-a0bb-87c4767e63d6</RequestId>
  </ResponseMetadata>
</GetOrderReferenceDetailsResponse>`

	detail := GetOrderReferenceDetailsResponse{}
	if err := xml.Unmarshal([]byte(resp), &detail); err != nil {
		t.Errorf("no error should get when unmarshal order reference detail, but got %v", err)
	}

	fmt.Printf("%#v", detail)
}

func TestAuthorizeResponse(t *testing.T) {
	resp := `<AuthorizeResponse xmlns="http://mws.amazonservices.com/schema/OffAmazonPayments/2013-01-01">
  <AuthorizeResult>
    <AuthorizationDetails>
      <AuthorizationAmount>
        <CurrencyCode>JPY</CurrencyCode>
        <Amount>134.00</Amount>
      </AuthorizationAmount>
      <CapturedAmount>
        <CurrencyCode>JPY</CurrencyCode>
        <Amount>0</Amount>
      </CapturedAmount>
      <ExpirationTimestamp>2018-04-01T06:21:12.317Z</ExpirationTimestamp>
      <IdList/>
      <SoftDecline>false</SoftDecline>
      <AuthorizationStatus>
        <LastUpdateTimestamp>2018-03-02T06:21:12.317Z</LastUpdateTimestamp>
        <State>Pending</State>
      </AuthorizationStatus>
      <AuthorizationFee>
        <CurrencyCode>JPY</CurrencyCode>
        <Amount>0.00</Amount>
      </AuthorizationFee>
      <CaptureNow>false</CaptureNow>
      <SellerAuthorizationNote/>
      <CreationTimestamp>2018-03-02T06:21:12.317Z</CreationTimestamp>
      <AmazonAuthorizationId>S03-0775027-3039214-A074087</AmazonAuthorizationId>
      <AuthorizationReferenceId>225</AuthorizationReferenceId>
    </AuthorizationDetails>
  </AuthorizeResult>
  <ResponseMetadata>
    <RequestId>a4e98146-b2f9-43b8-8ebf-364e1454b222</RequestId>
  </ResponseMetadata>
</AuthorizeResponse>`

	detail := AuthorizeResponse{}
	if err := xml.Unmarshal([]byte(resp), &detail); err != nil {
		t.Errorf("no error should get when unmarshal order reference detail, but got %v", err)
	}

	fmt.Printf("%#v", detail)
}

func TestErrorResponse(t *testing.T) {
	resp := `<ErrorResponse xmlns="http://mws.amazonservices.com/schema/OffAmazonPayments/2013-01-01">
	<Error>
	<Type>Sender</Type>
	<Code>DuplicateReferenceId</Code>
	<Message>Your Capture request could not be processed because an Authorize request with the ReferenceId 226 already exists.</Message>
	</Error>
	<RequestId>9547ad10-46da-4759-b37c-82ad02b0c1af</RequestId>
	</ErrorResponse>`

	detail := APIError{}
	if err := xml.Unmarshal([]byte(resp), &detail); err != nil {
		t.Errorf("no error should get when unmarshal order reference detail, but got %v", err)
	}

	if detail.Code != "DuplicateReferenceId" {
		t.Errorf("Error not decoded")
	}
}
