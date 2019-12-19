package champiris

import (
	"github.com/graphql-go/graphql"
	"github.com/kataras/iris/v12"
)

var Query RootType
var Mutation RootType

func init() {
	Query.New("Query", "搜尋&取得資料的相關命令")
	Mutation.New("Mutation", "主要用在建立、修改、刪除的相關命令")
}

type Ql struct {
	Ctx iris.Context
}

func (ql *Ql) Post() {
	params := ql.createParams()
	res := graphql.Do(params)
	_, _ = ql.Ctx.JSON(res)
}

func (ql *Ql) createParams() graphql.Params {
	opt := requestOptions{}
	_ = ql.Ctx.ReadJSON(&opt)
	return graphql.Params{
		Schema:         ql.newSchema(),
		RequestString:  opt.Query,
		VariableValues: opt.Variables,
		OperationName:  opt.OperationName,
	}
}

func (ql *Ql) newSchema() graphql.Schema {
	s, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    Query.Obj,
		Mutation: Mutation.Obj,
	})
	return s
}

type requestOptions struct {
	Query         string                 `json:"query" url:"query" schema:"query"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
	OperationName string                 `json:"operationName" url:"operationName" schema:"operationName"`
}
