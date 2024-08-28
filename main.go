package main

import (
	"errors"
	"fmt"
)

type Student struct {
	Name   string
	Height int
}

func (s *Student) Validate() error {

	if s.Name == "" {
		return errors.New("nama tidak boleh kosong")
	}

	if len(s.Name) <= 3 {
		panic("nama harus lebih besar dari 3")
	}

	if s.Height == 0 {
		return errors.New("tinggi tidak boleh kosong")
	}

	return nil

}

func main() {
	defer catchError()
	student := Student{Name: "dim", Height: 18}

	err := student.Validate()
	if err != nil {
		cetak(err.Error())
	} else {
		cetak("hello  " + student.Name)
	}

}


func cetak(text string) {
	fmt.Println(text)
}

func catchError() {
	err := recover()

	if err != nil {
		cetak("error nih : " + fmt.Sprint(err))
	}
}

