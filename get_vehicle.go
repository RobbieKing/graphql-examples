type package graphqlexamples

var client = graphql.Client

func ProcessVehicle(reg string) error{
	vehicle, err := graphql.GetVehicle(client, reg)
	if err != nil {
		return err
	}

	// Do some work
	return nil
}
