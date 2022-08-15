curl 'https://server/graphql' \ 
    -X POST \ 
	-H 'Content-Type: application/json' \
	--data-binary '{
		"query":
			"mutation {
				createVehicle(
					reg: "GOPHER2", 
					make: "GopherMobile", 
					odometer: 100
				) {
					reg
				}
			}"'
