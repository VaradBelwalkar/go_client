package requests

import (
	"fmt"
	"io/ioutil"
	"log"
	"bufio"
)

func Container_List(){

	request_path:="list/"
	//resp is of type map[string]interface{}
	// Here the object is derived from the JSON received from the response
	resp,err:= GET_Request(request_path)  

	if err!=nil {
		fmt.Prinln(err)
		return 
	}
	
	containerList:=resp["list"]	

	fmt.Println(containerList)
	
	
}

