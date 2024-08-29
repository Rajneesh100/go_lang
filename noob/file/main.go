package main

import (
	"fmt"
	"os"
)

func main() {

	// file, err := os.Create("a.txt")

	// if err != nil {
	// 	fmt.Println("got some error ")
	// } else {
	// 	fmt.Println("file created sucessfully:", file)
	// }

	// defer file.Close()

	// var content string = "hi !"
	// size, er := io.WriteString(file, content+"\n")
	// if er != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("type of byte :%T ,value: %d\n", size, size)
	//------------------------------------------------------------------------------------
	// file1, er := os.Open("a.txt")

	// if er != nil {
	// 	fmt.Println("got some error in opening ")
	// } else {
	// 	fmt.Println("file opened sucessfully:", file1)
	// }

	// defer file1.Close()

	// buffer := make([]byte, 1024)

	// for {

	// 	n, err := file1.Read(buffer)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("got some error in reading :", err)
	// 		return
	// 	}
	// 	fmt.Println(string(buffer[:n]))
	// }

	// file1.Close()

	//-----------------------------------------------------------------------------

	//depricated and it read all content at once so if file is large then can cause issue with memory
	// content, err := ioutil.ReadFile("a.txt")

	content, err := os.ReadFile("a.txt")
	if err != nil {
		fmt.Println("got some error")

	}
	fmt.Println(string(content))

}
