package tests

import (
	"errors"
	"git.championtek.com.tw/go/champiris"
	"github.com/graphql-go/graphql"
)

func addSchema() {
	champiris.Query.AddField(&graphql.Field{
		Name: "qq",
		Type: graphql.Int,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return 1, errors.New("WTF")
		},
	})
	champiris.Mutation.AddField(&graphql.Field{
		Name: "ff",
		Type: graphql.Int,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			type Member struct {
				ID int8 `json:"id"`
			}
			member := Member{}
			champiris.ToStruct(p.Args, &member)
			return member.ID, nil
		},
	})
}
