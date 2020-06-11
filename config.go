package champiris

type NetConfig struct {
	Protocol string `json:"Protocol" yaml:"Protocol"`
	Host     string `json:"Host" yaml:"Host"`
	Port     string `json:"Port" yaml:"Port"`
}
