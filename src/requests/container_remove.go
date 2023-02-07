package requests

import (
	"fmt"
	sh "github.com/VaradBelwalkar/go_client/session_handling"
)


func Container_Remove(containerName string){

	request_path:="/container/remove/"+containerName


	//resp is of type map[string]interface{}
	_,err:= sh.GET_Request(request_path)
	if err!=200{
		fmt.Println("something went wrong!")
	}

	//Here the backend is going to give JSON response containing info with 4 values,
	// 200 == OK
	// 300 == 


//Handle the response here
	
	
}

