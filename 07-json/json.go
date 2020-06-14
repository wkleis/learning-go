package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	marshal()
	unmarshal()

}

func marshal() {
	fmt.Println("Marshal")
	p1 := person{
		First: "John",
		Last:  "Developer",
		Age:   33,
	}
	p2 := person{
		First: "Marie",
		Last:  "Dubois",
		Age:   27,
	}
	people := []person{p1, p2}
	fmt.Println(people)
	jsonBytes, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)
}

func unmarshal() {
	//read given json into go struct
	fmt.Println("Unmarshal")
	jsonString := `[{"first":"Fred","last":"Miller","age":63},{"first":"Julio","last":"Martinez","age":57}]`
	fmt.Println(jsonString)
	bytes := []byte(jsonString)
	var persons []person
	err := json.Unmarshal(bytes, &persons)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(persons)
}
