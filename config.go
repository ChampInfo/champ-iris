package champiris

type NetConfig struct {
	Protocol        string  `json:"Protocol" yaml:"Protocol"`
	Host            string  `json:"Host" yaml:"Protocol"`
	Port            string  `json:"Port" yaml:"Protocol"`
	LoggerEnable    bool    `json:"LoggerEnable" yaml:"Protocol"`
	JWTEnable       bool    `json:"JWTEnable" yaml:"Protocol"`
}
