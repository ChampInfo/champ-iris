package tests

import (
	"errors"

	"github.com/graphql-go/graphql"
)

func addSchema() {
	Ql.Query.AddField(&graphql.Field{
		Name: "qq",
		Type: graphql.Int,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return 1, errors.New("WTF")
		},
	})
	Ql.Mutation.AddField(&graphql.Field{
		Name: "ff",
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Member",
			Fields: graphql.BindFields(Member{}),
		}),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			member := Member{}
			Ql.ToStruct(p.Args, &member)
			return member, nil
		},
	})
}

type Member struct {
	ID  int    `json:"id"`
	Map string `json:"map"`
}
