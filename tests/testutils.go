package tests

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
)

func addSchema() {
	Query.AddField(&graphql.Field{
		Name: "qq",
		Type: graphql.Int,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return 1, errors.New("WTF")
		},
	})
	Mutation.AddField(&graphql.Field{
		Name: "ff",
		Type: graphql.Int,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"Upload": &graphql.ArgumentConfig{
				Type: graphql.NewList(UploadScalar),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			type Member struct {
				ID int `json:"id"`
			}
			member := Member{}
			ToStruct(p.Args, &member)
			file := p.Args["Upload"].([]interface{})
			fmt.Println(file)
			return member.ID, nil
		},
	})
}
