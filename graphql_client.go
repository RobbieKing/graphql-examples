package graphqlexamples



type Client interface{
	Query(Request) ([]byte, error)
	Mutate(Request) error
}

const createVehicleMutation = `mutation (
	reg: String
){
	createVehicle(reg: $reg) {
		reg
	}
}
`

type Variables interface{}

type Request struct {
	Query     string    `json:"query"`
	Variables Variables `json:"variables"`
}

type CreateVehicleInput struct {
	Reg string `json:"reg"`
}


func CreateVehicle(client Client, reg string) error {
	req := Request{
		Variables: CreateVehicleInput{
			Reg: reg,
		},
		Query: createVehicleMutation,
	}

	if err := client.Mutate(req); err != nil {
		return err
	}

	return nil
}
