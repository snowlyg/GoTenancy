package amazonpay

import (
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Response Types
////////////////////////////////////////////////////////////////////////////////

// Profile user profile
type Profile struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

// OAuthResponseError oauth error response
type OAuthResponseError struct {
	ErrorDescription string `json:"error_description"`
	Error            string `json:"error"`
}

// ResponseMetadata respones meta data
type ResponseMetadata struct {
	RequestID string `xml:"RequestId"`
}

// SetOrderReferenceDetailsResult set order reference detail
type SetOrderReferenceDetailsResult struct {
	SetOrderReferenceDetailsResult struct {
		OrderReferenceDetails OrderReferenceDetails
	}
	ResponseMetadata ResponseMetadata
}

// GetOrderReferenceDetailsResponse details and current state of the Order Reference object.
type GetOrderReferenceDetailsResponse struct {
	GetOrderReferenceDetailsResult struct {
		OrderReferenceDetails OrderReferenceDetails
	}
	ResponseMetadata ResponseMetadata
}

// AuthorizeResponse authorize response
type AuthorizeResponse struct {
	AuthorizeResult struct {
		AuthorizationDetails AuthorizationDetails
	}
	ResponseMetadata ResponseMetadata
}

// GetAuthorizationDetailsResponse get authorization detail response
type GetAuthorizationDetailsResponse struct {
	GetAuthorizationDetailsResult struct {
		AuthorizationDetails
	}
	ResponseMetadata ResponseMetadata
}

// CaptureResponse capture response
type CaptureResponse struct {
	CaptureResult struct {
		CaptureDetails CaptureDetails
	}
	ResponseMetadata ResponseMetadata
}

// GetCaptureDetailsResponse get capture details response
type GetCaptureDetailsResponse struct {
	GetCaptureDetailsResult struct {
		CaptureDetails CaptureDetails
	}
	ResponseMetadata ResponseMetadata
}

// RefundResponse get refund response
type RefundResponse struct {
	RefundResult struct {
		RefundDetails RefundDetails
	}
	ResponseMetadata ResponseMetadata
}

// GetRefundDetailsResponse get refund detail response
type GetRefundDetailsResponse struct {
	GetRefundDetailsResult struct {
		RefundDetails RefundDetails
	}
	ResponseMetadata ResponseMetadata
}

////////////////////////////////////////////////////////////////////////////////
// Data Types
////////////////////////////////////////////////////////////////////////////////

// Address postal address information
type Address struct {
	Name                                     string
	AddressLine1, AddressLine2, AddressLine3 string
	City                                     string
	Country                                  string
	District                                 string
	StateOrRegion                            string
	PostalCode                               string
	CountryCode                              string
	Phone                                    string
}

// AuthorizationDetails details and status of the authorization object, including sales charge amount.
type AuthorizationDetails struct {
	AmazonAuthorizationID    string `xml:"AmazonAuthorizationId"`
	AuthorizationReferenceID string `xml:"AuthorizationReferenceId"`
	SellerAuthorizationNote  string
	AuthorizationAmount      Price
	CaptureAmount            Price
	AuthorizationFee         Price
	IDList                   []string `xml:"IdList"`
	CreationTimestamp        *time.Time
	ExpirationTimestamp      *time.Time
	AuthorizationStatus      Status
	SoftDecline              bool
	CaptureNow               bool
	SoftDescriptor           string
}

// Buyer buyer info
type Buyer struct {
	Name  string
	Email string
	Phone string
}

// CaptureDetails the details of the sales claim object and its current state.
type CaptureDetails struct {
	AmazonCaptureID    string `xml:"AmazonCaptureId"`
	CaptureReferenceID string `xml:"CaptureReferenceId"`
	SellerCaptureNote  string
	CaptureAmount      Price
	RefundAmount       Price
	CaptureFee         Price
	IDList             []string `xml:"IdList"`
	CreationTimestamp  *time.Time
	CaptureStatus      Status
	SoftDescriptor     string
}

// Constraint represents a mistake or error information of a Billing Agreement or Order Reference object.
type Constraint struct {
	ConstraintID string `xml:"ConstraintId"`
	Description  string
}

// Destination the address selected by the buyer via the address book widget.
type Destination struct {
	DestinationType     string
	PhysicalDestination Address
}

// OrderReferenceAttributes attribute of the Order Reference object specified by the seller.
type OrderReferenceAttributes struct {
	OrderTotal            OrderTotal
	PlatformID            string `xml:"PlatformId"`
	SellerNote            string
	SellerOrderAttributes SellerOrderAttributes
}

// OrderReferenceDetails details and current state of the Order Reference object.
type OrderReferenceDetails struct {
	AmazonOrderReferenceID string `xml:"AmazonOrderReferenceId"`
	Buyer                  Buyer
	OrderTotal             OrderTotal
	SellerNote             string
	PlatformID             string `xml:"PlatformId"`
	Destination            Destination
	ReleaseEnvironment     string
	SellerOrderAttributes  SellerOrderAttributes
	OrderReferenceStatus   OrderReferenceStatus
	Constraints            []Constraint
	CreationTimestamp      *time.Time
	ExpirationTimestamp    *time.Time
	IDList                 []string `xml:"IdList"`
}

// OrderReferenceStatus the current state of the Order Reference object.
type OrderReferenceStatus struct {
	State               string
	LastUpdateTimestamp *time.Time
	ReasonCode          string
	ReasonDescription   string
}

// OrderTotal total order amount presented in this Order Reference.
type OrderTotal struct {
	CurrencyCode string
	Amount       string
}

// Price currency type and amount.
type Price struct {
	Amount       string
	CurrencyCode string
}

// RefundDetails details and the current state of the refund object.
type RefundDetails struct {
	AmazonRefundID    string `xml:"AmazonRefundId"`
	RefundReferenceID string `xml:"RefundReferenceId"`
	SellerRefundNote  string
	RefundType        string
	RefundAmount      Price
	FeeRefunded       Price
	CreationTimestamp *time.Time
	RefundStatus      Status
	SoftDescriptor    string
}

// SellerOrderAttributes provides detailed information on the Order Reference object
type SellerOrderAttributes struct {
	SellerOrderID     string `xml:"SellerOrderId"`
	StoreName         string
	CustomInformation string
}

// Status represents the current status of authorization object, sales request object, refund object.
type Status struct {
	State               string
	LastUpdateTimestamp *time.Time
	ReasonCode          string
	ReasonDescription   string
}

// OrderReferenceNotification order reference notification data
type OrderReferenceNotification struct {
	OrderReference OrderReferenceDetails
}

// AuthorizationNotification authorization notification data
type AuthorizationNotification struct {
	AuthorizationDetails AuthorizationDetails
}

// CaptureNotification capture notification
type CaptureNotification struct {
	CaptureDetails CaptureDetails
}

// RefundNotification refund notification
type RefundNotification struct {
	RefundDetails RefundDetails
}
