# Configuration
```go
type NetworkSetup struct {
	Mask   int    `json:"mask" yaml:"mask" mapstructure:"mask"`
	Subnet string `json:"subnet" yaml:"subnet" mapstructure:"subnet"`
}

type V4GatewayConfig struct {
	IpHeaderName string         `json:"ip-header-name" yaml:"ip-header-name" mapstructure:"ip-header-name"`
	Networks     []NetworkSetup `json:"networks" yaml:"networks" mapstructure:"networks"`
}
```
## Example
```yaml
ip-header-name: "X-Real-IP"
networks:
  - mask: 24
    subnet: "192.168.1.0"
  - mask: 16
    subnet: "10.0.0.0"
``` 