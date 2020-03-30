package elklogger

type GraphqlBody struct {
	Query         string      `json:"query"`
	Variables     interface{} `json:"variables"`
	OperationName string      `json:"operationName"`
}
