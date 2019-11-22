package champiris

import (
	"testing"

	"github.com/graphql-go/graphql"
)

func TestAPI_NewService(t *testing.T) {
	var server API
	_ = server.NewService(NetConfig{
		Port: "8080",
	})
	server.app.Logger().SetLevel("debug")
	addSchema()
	_ = server.Run()
}

func addSchema() {
	Query.AddField(&graphql.Field{
		Name: "qq",
		Type: graphql.Int,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return 1, nil
		},
	})
	Mutation.AddField(&graphql.Field{
		Name: "ff",
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			i2 := p.Args["id"].(int)
			return string(i2), nil
		},
	})
}
