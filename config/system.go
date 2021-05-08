package config

type System struct {
	Level     string `mapstructure:"level" json:"level" yaml:"level"`
	Env       string `mapstructure:"env" json:"env" yaml:"env"`
	Addr      int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType    string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	CacheType string `mapstructure:"cache-type" json:"cacheType" yaml:"cache-type"`
}
