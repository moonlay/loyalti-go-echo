package graphQL

import (
	"github.com/graphql-go/graphql"
)

// Root holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

type Page struct {
	page int `json:"null"`
}

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot() *Root {
	// Create a resolver holding our database. Resolver can be found in resolvers.go

	// Create a new Root that describes our base query set up. In this
	// example we have a user query that takes one argument called name
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"songs": &graphql.Field{
						// Slice of User type which can be found in types.go
						Type:    graphql.NewList(songType),
						Resolve: SongResolver,
					},
					"merchant": &graphql.Field{
						Type: graphql.NewList(merchantType),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"email": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: MerchantResolver,
					},
					"category": &graphql.Field{
						Type: graphql.NewList(categoryType),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: MerchantCategoryResolver,
					},
					"cardtype": &graphql.Field{
						Type: graphql.NewList(cardType),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: MerchantCardTypeResolver,
					},
					//"cardtier": &graphql.Field{
					//	Type: graphql.NewList(card),
					//	Args: graphql.FieldConfigArgument{
					//		"program_id": &graphql.ArgumentConfig{
					//			Type: graphql.Int,
					//		},
					//	},
					//	Resolve: MerchantCardMemberResolver,
					//},
					"socialmedia": &graphql.Field{
						Type: graphql.NewList(socialmediaType),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: SocialMediaResolver,
					},

					"province": &graphql.Field{
						Type: graphql.NewList(provinceType),
						Args:graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: ProvinceResolver,
					},

					"city": &graphql.Field{
						Type: graphql.NewList(cityType),
						Args:graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: CityResolver,
					},

					"program": &graphql.Field{
						Type: graphql.NewList(programType),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"category": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: ProgramResolver,
					},
					"special": &graphql.Field{
						Type: graphql.NewList(specialprogramType),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"category": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: SpecialProgramResolver,
					},
					"outlet": &graphql.Field{
						Type: graphql.NewList(outletType),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"email": &graphql.ArgumentConfig{
									Type: graphql.String,
							},
						},
						Resolve: OutletResolver,
					},
					"employee": &graphql.Field{
						Type: graphql.NewList(employeeType),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: EmployeeResolver,
					},
					"totalpoint": &graphql.Field{
						Type: graphql.NewList(totalpointType),
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"pay": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"pin": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"outletid": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"cardtype": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: TotalPointResolver,
					},
					"totalchop": &graphql.Field{
						Type: graphql.NewList(totalchopType),
						Args: graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"pay": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"pin": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"outletid": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"cardtype": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: TotalChopResolver,
					},
					"merchanttransaction": &graphql.Field{
						Type: graphql.NewList(transactionType),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"outletid": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: TransactionResolver,
					},
					"card": &graphql.Field{
						Type: graphql.NewList(card),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"card_type": &graphql.ArgumentConfig{
									Type:graphql.String,
							},
						},
						Resolve: CardResolver,
					},
					"voucher": &graphql.Field{
						Type:graphql.NewList(voucher),
						Args:graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"merchant_id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve:VoucherResolver,
					},
				},
			},
		),
	}
	return &root
}
