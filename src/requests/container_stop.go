package requests

import (
	"fmt"
	"io/ioutil"
	"log"
	"bufio"
	"golang.org/x/crypto/ssh"
	sh "github.com/VaradBelwalkar/go_client/session_handling"
)


func Container_stop(containerName string){

	request_path:="stop/"+containerName


	//resp is of type map[string]interface{}
	resp,err := sh.GET_Request(request_path)

	if err!=nil{
		fmt.Println("Handle the error efficiently!")}
return
	
}

