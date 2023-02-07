package requests

import (
	"fmt"
	sh "github.com/VaradBelwalkar/go_client/session_handling"
)


func Container_stop(containerName string){

	request_path:="/container/stop/"+containerName


	//resp is of type map[string]interface{}
	_,err := sh.GET_Request(request_path)

	if err!=200{
		fmt.Println("Handle the error efficiently!")}
return
	
}

