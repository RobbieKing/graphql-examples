package graphqlexamples

import (
	"context"

	"github.com/Khan/genqlient/graphql"

	"github.com/RobbieKing/graphql-examples/gen"
)

type Vehicle struct {
	Reg string `json:"reg"`
}

type Client struct {
	genQlient              graphql.Client
}

func (c *Client) CreateVehicle(ctx context.Context, vehicle Vehicle) error {
	_, err := gen.CreateVehicle(ctx, c.genQlient, vehicle.Reg)
	return err
}
