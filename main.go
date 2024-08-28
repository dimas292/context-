package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type Data struct {
	Method string
	Value  string
}

func main() {

	ctx := context.Background()

	dataA := Data{
		Method: "GET",
		Value: "Mengambil data dari context A",
	}

	ctxLevel1_A := context.WithValue(ctx, "data", dataA)

	dataB := Data{
		Method: "POST",
		Value: "Menambah data user pada context B",
	}

	ctxLevel1_B := context.WithValue(ctx, "data", dataB)

	processA(ctxLevel1_A)
	processB(ctxLevel1_B)
	

}

func processA(ctx context.Context){
	var data Data

	fmt.Println("proccess context level 1 A")

	dataCtx := ctx.Value("data")

	dataCtxByte, err := json.Marshal(dataCtx)
	if err != nil {
		fmt.Println("eror parsing")
		return
	}

	err = json.Unmarshal(dataCtxByte, &data)

	fmt.Println("try proccess with method: ", data.Method)
	fmt.Println(data.Value)
	
}
func processB(ctx context.Context){

	var data Data

	fmt.Println("proccess context level 1 A")

	dataCtx := ctx.Value("data")

	dataCtxByte, err := json.Marshal(dataCtx)
	if err != nil {
		fmt.Println("error parsing")
		return
	}

	err = json.Unmarshal(dataCtxByte, &data)

	fmt.Println("try proccess with method: ", data.Method)
	fmt.Println(data.Value)

}
