package main

import (
	"fmt"
	"net/url"
)

func main() {

	myurl := "https://trabii.azurewebsites.net/orgs/get_org_details/1"
	fmt.Printf("typr of utl :%T\n", myurl)

	parseurl, er := url.Parse(myurl)

	if er != nil {
		fmt.Println("got some issue in parsing")
		return
	}

	fmt.Printf("type of parseurl : %T\n", parseurl)
	fmt.Println(parseurl)

	fmt.Println("Scheme:", parseurl.Scheme)
	fmt.Println("Host: ", parseurl.Host)
	fmt.Println("path:", parseurl.Path)
	fmt.Println("raw query :", parseurl.RawQuery)
	parseurl.Path = "/newPath"
	parseurl.RawQuery = "username=rajneesh"

	fmt.Println("New url:", parseurl)

}
