package graphqlexamples

import (
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
)


const createVehicleMutation = `mutation (
		reg: String
	){
		createVehicle(reg: $reg) {
			reg
		}
	}
`

type request struct {
	Variables map[string]interface{} `json:"variables"`
	Query     string                 `json:"query"`
}

func createVehicle(reg string) error {
	req := request{
		Variables: map[string]interface{}{
			"reg": reg,
		},
		Query: createVehicleMutation,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}


	resp, err := http.Post("http://server/graphql", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error performing graphql request: %d", resp.StatusCode)
	}

	return nil
}
