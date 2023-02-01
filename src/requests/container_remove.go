package requests

import (
	"fmt"
	"io/ioutil"
	"log"
	"bufio"
)


func Container_Run(containerName string){

	request_path:="remove/"+containerName


	//resp is of type map[string]interface{}
	resp := GET_Request(request_path)

//Handle more things here
	
	
}

