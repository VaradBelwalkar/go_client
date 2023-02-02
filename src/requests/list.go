package requests

import (
	"fmt"
	"io/ioutil"
	"log"
	"bufio"
	"golang.org/x/crypto/ssh"
)


func List(what string){

	request_path:="list/"+what


	//resp is of type map[string]interface{}
	resp := GET_Request(request_path)

	list:=resp["list"]	

	fmt.Println(list)
	
	
}

