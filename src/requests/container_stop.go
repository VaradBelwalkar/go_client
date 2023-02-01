package requests

import (
	"fmt"
	"io/ioutil"
	"log"
	"bufio"
	"golang.org/x/crypto/ssh"
)


func Container_stop(containerName string){

	request_path:="stop/"+containerName


	//resp is of type map[string]interface{}
	resp,err := GET_Request(request_path)

	iferr!=nil{
		fmt.Println("Handle the error efficiently!")
		return 
	}

	
}

