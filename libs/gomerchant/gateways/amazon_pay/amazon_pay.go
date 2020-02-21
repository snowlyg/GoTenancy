package amazon_pay

// AmazonPay amazon pay
type AmazonPay struct {
	*Config
}

// Config amazon pay configuration
type Config struct {
	MerchantID   string
	AccessKey    string
	SecretKey    string
	Region       string
	CurrencyCode string

	ProductionMode bool
}

// New initialize amazon pay
func New(config *Config) *AmazonPay {
	if config == nil {
		config = &Config{}
	}

	return &AmazonPay{Config: config}
}
