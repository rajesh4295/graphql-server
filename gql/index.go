package gql

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/rajesh4295/graphql-server/database"
	"github.com/rajesh4295/graphql-server/models"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"walletId": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					user, err := database.GetUserById(p.Args["id"].(string))
					if err != nil {
						return nil, err
					}
					return user, nil
				},
			},
			"users": &graphql.Field{
				Type: graphql.NewList(userType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					users, err := database.GetUsers()
					if err != nil {
						return nil, err
					}
					return users, nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"userAdd": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"walletId": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					allUsers, _ := database.GetUsers()
					newId := fmt.Sprintf("%d", len(allUsers)+1)
					newUser := &models.User{Id: newId, Name: (p.Args["name"]).(string), WalletId: (p.Args["walletId"]).(string)}
					user, err := database.AddUser(newUser)
					if err != nil {
						return nil, err
					}
					return user, nil
				},
			},
			"userDelById": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					userId := p.Args["id"].(string)
					existingUser, err := database.GetUserById(userId)
					if err != nil {
						return nil, err
					}
					user, err := database.RemoveUserById(existingUser)
					if err != nil {
						return nil, err
					}
					return user, nil
				},
			},
			"userUpdateById": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"name": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
					"walletId": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					existingUser := &models.User{Id: (p.Args["id"]).(string), Name: (p.Args["name"]).(string), WalletId: (p.Args["walletId"]).(string)}
					user, err := database.UpdateUserById(existingUser)
					if err != nil {
						return nil, err
					}
					return user, nil
				},
			},
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
