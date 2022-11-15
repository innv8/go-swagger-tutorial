package entities

// swagger:model
type Person struct {
	// the id of the user
	//
	// required: true
	// min: 1
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}
