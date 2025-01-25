package gateway

type NetworkSetup struct {
	Mask   int    `json:"mask" yaml:"mask" mapstructure:"mask"`
	Subnet string `json:"subnet" yaml:"subnet" mapstructure:"subnet"`
}

type V4GatewayConfig struct {
	IpHeaderName string          `json:"ip-header-name" yaml:"ip-header-name" mapstructure:"ip-header-name"`
	Networks     []*NetworkSetup `json:"networks" yaml:"networks" mapstructure:"networks"`
}

var conf *V4GatewayConfig

func UploadGatewayConfig(config *V4GatewayConfig) error {
	// Save config to your preferred storage
	conf = config
	return nil
}
