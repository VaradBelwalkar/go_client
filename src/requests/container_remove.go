package requests

import (
	"fmt"
	"io/ioutil"
	"log"
	"bufio"
)


func Container_Remove(containerName string){

	request_path:="remove/"+containerName


	//resp is of type map[string]interface{}
	resp := GET_Request(request_path)

	//Here the backend is going to give JSON response containing info with 4 values,
	// 200 == OK
	// 300 == 


//Handle the response here
	
	
}

