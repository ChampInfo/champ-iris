package elklogger

type GraphqlBody struct {
	Query           string  `json:"query"`
	Variables       string  `json:"variables"`
	OperationName   string  `json:"operationName"`
}