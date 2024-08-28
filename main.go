package main

import (
	"encoding/json"
	"fmt"
)

type User struct{
	Name string
	Class string
}

func main(){
	users := []User{
		{Name: "Dimas", Class: "6A"},
		{Name: "Java", Class: "7A"},
		{Name: "Nodejs", Class: "5A"},
	}

	userJson, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("bentuk json : ", string(userJson))
	

}