package champiris

type IrisNetWork struct {
	Protocol string `json:"Protocol"`
	Host     string `json:"Host"`
	Port     string `json:"Port"`
	WebPath  bool   `json:"WebPath"`
	ELK      ELK    `json:"elk"`
}

type ELK struct {
	Host      ElkHost    `json:"Host"`
	Index     string     `json:"Index"`
	BasicAuth Auth       `json:"BasicAuth"`
	Mapping   ElkMapping `json:"Mapping"`
}

type Auth struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type ElkHost struct {
	URL  string `json:"url"`
	PORT string `json:"port"`
}

type ElkMapping struct {
	Settings Settings `json:"settings"`
	Mappings Mappings `json:"mappings"`
}

type Settings struct {
	NumberOfShards   int `json:"number_of_shards"`
	NumberOfReplicas int `json:"number_of_replicas"`
}

type Mappings struct {
	Properties Properties `json:"properties"`
}

type Properties struct {
	Service string `json:"service"`
	IP      string `json:"request_ip"`
	Status  string `json:"status"`
	Method  string `json:"method"`
	Path    string `json:"path"`
	Tags    string `json:"tags"`
	Created string `json:"created"`
	Remark  string `json:"remark"`
}
