package requests

import (
	"fmt"
	sh "github.com/VaradBelwalkar/go_client/session_handling"
)

func Container_List(){

	request_path:="/container/list/containers"
	//resp is of type map[string]interface{}
	// Here the object is derived from the JSON received from the response
	resp,err:= sh.GET_Request(request_path)  

	if err!=200 {
		if err==500{
			fmt.Println("Server error!")
		return
		}
		 if err==404{
			fmt.Println("something")
		 }

	}
	var resultArray []string
	for _,v:=range resp{
		resultArray=append(resultArray,v.(string))
	}
	for _, k := range resultArray { 
		fmt.Println(k)
		
	}	
	
}

