package champiris

type NetConfig struct {
	Protocol        string  `json:"Protocol" yaml:"Protocol"`
	Host            string  `json:"Host" yaml:"Host"`
	Port            string  `json:"Port" yaml:"Port"`
	LoggerEnable    bool    `json:"LoggerEnable" yaml:"LoggerEnable"`
	JWTEnable       bool    `json:"JWTEnable" yaml:"JWTEnable"`
}
