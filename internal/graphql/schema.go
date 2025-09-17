package graphql

import (
	_ "fmt"

	"github.com/graphql-go/graphql"
	"github.com/dopaemon/KetQuaXoSo/internal/configs"
	"github.com/dopaemon/KetQuaXoSo/internal/rss"
	"github.com/dopaemon/KetQuaXoSo/utils"
)

var checkResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "CheckResponse",
	Fields: graphql.Fields{
		"province": &graphql.Field{Type: graphql.String},
		"results":  &graphql.Field{Type: graphql.NewList(graphql.String)},
		"error":    &graphql.Field{Type: graphql.String},
	},
})

var ticketResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TicketResponse",
	Fields: graphql.Fields{
		"province": &graphql.Field{Type: graphql.String},
		"date":     &graphql.Field{Type: graphql.String},
		"input":    &graphql.Field{Type: graphql.String},
		"prize":    &graphql.Field{Type: graphql.String},
		"match":    &graphql.Field{Type: graphql.String},
		"error":    &graphql.Field{Type: graphql.String},
	},
})

func GetSchema() (graphql.Schema, error) {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"provinces": &graphql.Field{
				Type: graphql.NewList(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return configs.Provinces, nil
				},
			},
			"checkLottery": &graphql.Field{
				Type: checkResponseType,
				Args: graphql.FieldConfigArgument{
					"province": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					province := p.Args["province"].(string)
					url, _ := rss.Sources(province)
					if url == "" {
						return map[string]interface{}{
							"province": province,
							"error":    "Unknown province: " + province,
						}, nil
					}

					data, err := rss.Fetch(url)
					if err != nil {
						return map[string]interface{}{
							"province": province,
							"error":    err.Error(),
						}, nil
					}

					results, err := rss.Parse(data)
					if err != nil {
						return map[string]interface{}{
							"province": province,
							"error":    err.Error(),
						}, nil
					}

					return map[string]interface{}{
						"province": province,
						"results":  results,
					}, nil
				},
			},
			"checkTicket": &graphql.Field{
				Type: ticketResponseType,
				Args: graphql.FieldConfigArgument{
					"province": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"date":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"number":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					province := p.Args["province"].(string)
					date := p.Args["date"].(string)
					number := p.Args["number"].(string)

					url, _ := rss.Sources(province)
					if url == "" {
						return map[string]interface{}{
							"province": province,
							"date":     date,
							"input":    number,
							"error":    "Unknown province: " + province,
						}, nil
					}

					data, err := rss.Fetch(url)
					if err != nil {
						return map[string]interface{}{
							"province": province,
							"date":     date,
							"input":    number,
							"error":    err.Error(),
						}, nil
					}

					results, err := rss.Parse(data)
					if err != nil {
						return map[string]interface{}{
							"province": province,
							"date":     date,
							"input":    number,
							"error":    err.Error(),
						}, nil
					}

					prize, match := utils.CheckWinningNumber(results, date, number)

					return map[string]interface{}{
						"province": province,
						"date":     date,
						"input":    number,
						"prize":    prize,
						"match":    match,
					}, nil
				},
			},
			"dates": &graphql.Field{
				Type: graphql.NewList(graphql.String),
				Args: graphql.FieldConfigArgument{
					"province": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					province := p.Args["province"].(string)

					url, _ := rss.Sources(province)
					if url == "" {
						return nil, nil
					}

					data, err := rss.Fetch(url)
					if err != nil {
						return nil, err
					}

					results, err := rss.Parse(data)
					if err != nil {
						return nil, err
					}

					var dates []string
					for _, v := range results {
						dates = append(dates, v.Date)
					}

					return dates, nil
				},
			},

		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
}
