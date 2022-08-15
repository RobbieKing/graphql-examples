package graphqlexamples


func TestCreateVehiclesV2(t *testing.T) {
	// Given
	testMock := graphqlMock{}

	vehicle := Vehicle{
		Reg: "GOPHER1",
	}
	
	testMock.On("GetVehicle").Return(vehicle, nil)

	// When
	err := CreateVehiclesV2()
	
	// Then
	assert.NoError(t, err)
	testMock.AssertExpectations(t)
}
