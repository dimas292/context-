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
	var users []User

	fmt.Println("users saat ini : ", len(users))
	
	dataDariFrontEnd := `{"Name": "Dimas", "Class": "B"}`
	
	data := []byte(dataDariFrontEnd)
	
	var user User
	
	err := json.Unmarshal(data, &user)
	
	if err != nil {
		fmt.Println("error saat encoding")
	}
	
	users = append(users, user)
	
	fmt.Println("users saat ini : ", len(users))
	fmt.Println("data user -> : " , users)
}