package tests

import (
	"encoding/json"
	"mime/multipart"
	"strconv"
	"strings"

	"git.championtek.com.tw/go/champiris"
	"github.com/graphql-go/graphql"
	"github.com/kataras/iris/v12"
)

var Query champiris.RootType
var Mutation champiris.RootType

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
	opt := ql.getRequestOptions()
	return graphql.Params{
		Schema:         ql.newSchema(),
		RequestString:  opt.Query,
		VariableValues: opt.Variables,
		OperationName:  opt.OperationName,
	}
}

func (ql *Ql) getRequestOptions() *requestOptions {
	r := ql.Ctx.Request()
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		var opts requestOptions
		_ = ql.Ctx.ReadJSON(&opts)
		return &opts
	}

	operationsParam := r.FormValue("operations")
	var opts requestOptions
	if err := json.Unmarshal([]byte(operationsParam), &opts); err != nil {
		return &requestOptions{}
	}

	mapParam := r.FormValue("map")
	mapValues := make(map[string][]string)
	if len(mapParam) != 0 {
		if err := json.Unmarshal([]byte(mapParam), &mapValues); err != nil {
			return &requestOptions{}
		}
	}

	variables := opts

	for key, value := range mapValues {
		for _, v := range value {
			if file, header, err := r.FormFile(key); err == nil {

				// Now set the path in ther variables
				var node interface{} = variables

				parts := strings.Split(v, ".")
				last := parts[len(parts)-1]

				for _, vv := range parts[:len(parts)-1] {
					switch node.(type) {
					case requestOptions:
						if vv == "variables" {
							node = opts.Variables
						} else {
							return &requestOptions{}
						}
					case map[string]interface{}:
						node = node.(map[string]interface{})[vv]
					case []interface{}:
						if idx, err := strconv.ParseInt(vv, 10, 64); err == nil {
							node = node.([]interface{})[idx]
						} else {
							return &requestOptions{}
						}
					default:
						return &requestOptions{}
					}
				}

				data := &MultipartFile{File: file, Header: header}

				switch node.(type) {
				case map[string]interface{}:
					node.(map[string]interface{})[last] = data
				case []interface{}:
					if idx, err := strconv.ParseInt(last, 10, 64); err == nil {
						node.([]interface{})[idx] = data
					} else {
						return &requestOptions{}
					}
				default:
					return &requestOptions{}
				}
			}
		}
	}
	return &opts
}

func (ql *Ql) newSchema() graphql.Schema {
	s, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    Query.Obj,
		Mutation: Mutation.Obj,
	})
	return s
}

type MultipartFile struct {
	File   multipart.File
	Header *multipart.FileHeader
}

var UploadScalar = graphql.NewScalar(graphql.ScalarConfig{
	Name: "Upload",
	ParseValue: func(value interface{}) interface{} {
		if v, ok := value.(*MultipartFile); ok {
			return v
		}
		return nil
	},
})

type requestOptions struct {
	Query         string                 `json:"query" url:"query" schema:"query"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
	OperationName string                 `json:"operationName" url:"operationName" schema:"operationName"`
}
