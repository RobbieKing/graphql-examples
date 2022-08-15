package graphqlexamples

func TestGetVehicle(t *testing.T) {
	// Given
	testMock := graphqlMock{}

	vehicle := &graptql.GetVehicleResult{
		Data: graptql.GetVehicleData{
			GetVehicle: Vehicle{Reg: "GOPHER1"},
		},
	}

	jsonData, err := json.Marshal(vehicle)
	
	testMock.On("Query").Return(jsonData, nil)

	// When
	err := ProcessVehicle()

	// Then
	assert.NoError(t, err)
	testMock.AssertExpectations(t)
}
