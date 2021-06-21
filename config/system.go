package config

type System struct {
	Level       string `mapstructure:"level" json:"level" yaml:"level"`
	Env         string `mapstructure:"env" json:"env" yaml:"env"`
	Addr        int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType      string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	OssType     string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"` // Oss类型
	CacheType   string `mapstructure:"cache-type" json:"cacheType" yaml:"cache-type"`
	AdminPreix  string `mapstructure:"admin-preix" json:"adminPreix" yaml:"admin-preix"`
	ClientPreix string `mapstructure:"client-preix" json:"clientPreix" yaml:"client-preix"`
}
