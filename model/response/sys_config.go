package response

type SysConfig struct {
	TenancyResponse
	ConfigName string `json:"configName"  binding:"required"` // 字段名称
	ConfigKey  string `json:"configKey"  binding:"required"`  // 字段 key
	ConfigType string `json:"configType"  binding:"required"` // 配置类型
	ConfigRule string `json:"configRule"`                     // 规则
	Required   int    `json:"required"`                       // 必填
	Info       string `json:"info"`                           // 配置说明
	Sort       uint16 `json:"sort"`                           // 排序
	TypeName   string `json:"typeName"`                       // 配置类型
	Value      string `json:"value"`                          // 配置类型
}
