package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("learing api")
	// res, er := http.Get("https://trabii.azurewebsites.net/event/details/0")

	res, er := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if er != nil {
		fmt.Println("Didn't got api req")
		return
	}

	defer res.Body.Close()

	// fmt.Println(res)
	fmt.Printf("type of res: %T\n", res)
	data, er1 := ioutil.ReadAll(res.Body) // here return data is array of bytes

	if er1 != nil {
		fmt.Println("got some errror in reading response :", er1)
		return
	}
	fmt.Println(string(data))

}
