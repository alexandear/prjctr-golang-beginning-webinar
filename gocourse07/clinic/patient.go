package clinic

import (
	"fmt"
	"math/rand/v2"
)

type Patient struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	BloodType string `json:"blood_type"`
	Data      string `json:"data"`
}

var (
	names      = []string{"John", "Jane", "Bob", "Alice", "Mike", "Emily", "David", "Sarah"}
	bloodTypes = []string{"A", "B", "AB", "O"}
)

func GenerateRandomPatient() Patient {
	id := fmt.Sprintf("P%03d", rand.N(1000))
	name := names[rand.N(len(names))]
	age := rand.N(80) + 1
	bloodType := bloodTypes[rand.N(len(bloodTypes))]

	return Patient{
		ID:        id,
		Name:      name,
		Age:       age,
		BloodType: bloodType,
	}
}
