package requests

import (
	"fmt"
	"io/ioutil"
	"log"
	"bufio"
	sh "github.com/VaradBelwalkar/go_client/session_handling"
)

func Container_List(){

	request_path:="list/"
	//resp is of type map[string]interface{}
	// Here the object is derived from the JSON received from the response
	resp,err:= sh.GET_Request(request_path)  

	if err!=nil {
		fmt.Prinln(err)
		return 
	}
	
	containerList:=resp.([]string)	
	for _, k := range containerList { 
		fmt.Println(k)
		
	}
	fmt.Println(containerList)
	
	
}

