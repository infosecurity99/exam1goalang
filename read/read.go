package read

import (
	"encoding/json"
	"os"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Cash      int    `json:"cash"`
	Basket    Basket `json:"basket"`
}

type Basket struct {
	Id       string    `json:"id"`
	Products []Product `json:"products"`
	Total    int       `json:"total"`
}

type Product struct {
	Id       string `json:"id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func OpenJson() ([]User, error) {
	users := []User{}

	file, err := os.Open("dataJson/first_exam.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&users); err != nil {
		return nil, err
	}
	return users, nil
}
