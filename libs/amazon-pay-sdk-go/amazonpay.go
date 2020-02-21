package amazonpay

// MWSServiceURLs mwsservice urls
var MWSServiceURLs = map[string]string{
	"eu": "mws-eu.amazonservices.com",
	"na": "mws.amazonservices.com",
	"jp": "mws.amazonservices.jp",
}

var (
	// OAuthEndpoint oauth endpoint
	OAuthEndpoint = "https://api.amazon.com"
	// OAuthSandboxEndpoint oauth endpoint for sandbox env
	OAuthSandboxEndpoint = "https://api.sandbox.amazon.com"
)

// RegionMappings region mapping
var RegionMappings = map[string]string{
	"de": "eu",
	"uk": "eu",
	"us": "na",
	"jp": "jp",
}

type AmazonPayService interface {
	GetProfile(token string) (profile Profile, err error)
	SetOrderReferenceDetails(orderReferenceID string, attrs OrderReferenceAttributes) (result SetOrderReferenceDetailsResult, err error)
	ConfirmOrderReference(orderReferenceID string) error
	GetOrderReferenceDetails(orderReferenceID string, addressToken string) (result GetOrderReferenceDetailsResponse, err error)
	Authorize(orderReferenceID string, authorizationReferenceID string, amount Price, input AuthorizeInput) (result AuthorizeResponse, err error)
	GetAuthorizationDetails(authorizationID string) (result GetAuthorizationDetailsResponse, err error)
	CloseAuthorization(authorizationID string, closureReason string) error
	Capture(authorizationID string, captureReferenceID string, captureAmount Price, input CaptureInput) (result CaptureResponse, err error)
	GetCaptureDetails(captureID string) (result GetCaptureDetailsResponse, err error)
	CloseOrderReference(orderReferenceID string, closureReason string) error
	CancelOrderReference(orderReferenceID string, reason string) error
	Refund(captureID string, refundReferenceID string, refundAmount Price, input RefundInput) (result RefundResponse, err error)
	GetRefundDetails(refundID string) (result GetRefundDetailsResponse, err error)
	Post(params Params, response interface{}) error
	Sign(message string) string
}

// AmazonPay amazon pay
type AmazonPay struct {
	*Config
}

// Config Amazon Pay Config
//
// Note on 'PlatformID':
//   According to AmazonPay docs, 'PlatformID' is the MerchatID, which is also referenced as SellerID
//   https://developer.amazon.com/ja/docs/amazon-pay-api/setorderattributes.html
//
//   But, in real life business, AmazonPay vendor may provide a different PlatformID for tracking incentives
//   So you can set this seperately
type Config struct {
	MerchantID    string
	AccessKey     string
	SecretKey     string
	Sandbox       bool
	Region        string
	CurrencyCode  string
	Endpoint      string
	OAuthEndpoint string
	ModePath      string
	APIVersion    string
	PlatformID    string
}

// New initialize amazon pay
func New(config *Config) AmazonPayService {
	if config == nil {
		config = &Config{}
	}

	if region, ok := RegionMappings[config.Region]; ok {
		config.Region = region
	}

	if config.Region == "" {
		config.Region = "na"
	}

	if config.Endpoint == "" {
		config.Endpoint = MWSServiceURLs[config.Region]
	}

	if config.APIVersion == "" {
		config.APIVersion = "2013-01-01"
	}

	if config.ModePath == "" {
		if config.Sandbox {
			config.ModePath = "OffAmazonPayments_Sandbox"
		} else {
			config.ModePath = "OffAmazonPayments"
		}
	}

	if config.OAuthEndpoint == "" {
		if config.Sandbox {
			config.OAuthEndpoint = OAuthSandboxEndpoint
		} else {
			config.OAuthEndpoint = OAuthEndpoint
		}
	}

	return &AmazonPay{
		Config: config,
	}
}
