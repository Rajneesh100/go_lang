package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getresponse() {

	res, er := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	//using error
	if er != nil {
		fmt.Println("issue in get response : ", er)
		return
	}
	// we can check with this as well
	if res.StatusCode != http.StatusOK {
		fmt.Println("got error in response:", res.StatusCode)
		return
	}
	//free the resources at end
	defer res.Body.Close()

	// data, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(string(data))

	var todo Todo
	er1 := json.NewDecoder(res.Body).Decode(&todo)
	if er1 != nil {
		fmt.Println("got error in copying :", er1)
		return
	}

	fmt.Println(todo)
	fmt.Println(todo.Title)
}
func postresponse() {

	url := "https://jsonplaceholder.typicode.com/todos/"
	var todo Todo = Todo{UserId: 33, Id: 89, Title: "hasta la vista", Completed: true}

	//convert todo struct to json
	jsondata, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("issue in json conversion :", err)
		return
	}
	// json data to string
	jsonString := string(jsondata)

	jsonReader := strings.NewReader(jsonString)
	res, er := http.Post(url, "application/json", jsonReader)

	if er != nil {
		fmt.Println("error in post")
		return
	}
	defer res.Body.Close()
	fmt.Println("Respons staus :", res.Status)
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))
}
func putrequest() {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	var todo Todo = Todo{UserId: 33, Id: 89, Title: "hasta la vista", Completed: true}

	//convert todo struct to json
	jsondata, err1 := json.Marshal(todo)
	if err1 != nil {
		fmt.Println("issue in json conversion :", err1)
		return
	}
	// json data to string
	jsonString := string(jsondata)

	//string data to reader
	jsonReader := strings.NewReader(jsonString)

	// create put request
	req, er := http.NewRequest(http.MethodPut, url, jsonReader)

	if er != nil {
		fmt.Println("error in put")
		return
	}

	req.Header.Set("Content-type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error in put :", err)
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))
	fmt.Println("Respons staus :", res.Status)
}
func deletereqest() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// create delete request
	req, er := http.NewRequest(http.MethodDelete, url, nil)

	if er != nil {
		fmt.Println("error in put")
		return
	}

	//send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error in put :", err)
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))
	fmt.Println("Respons staus :", res.Status)
}

func main() {
	// getresponse()
	// postresponse()
	// putrequest()
	deletereqest()
}
