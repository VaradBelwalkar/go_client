package session_handling

import (

	"fmt"

	"net/http"

	//"github.com/VaradBelwalkar/go_client"
)



//Handles the reponses from the server with default status codes
//It will handle and print appropriate response to the user so no handling in the client code
func Handle_resp_err(httpResponse *http.Response) bool{

	if httpResponse.StatusCode == 403{
		fmt.Println("\nWrong password!\tPlease update your password by running config edit\n")
		return true
	}else if httpResponse.StatusCode == 208{
		fmt.Println("User doesn't exist!Please create account by running \"setup\"or update your credentials by running \"config edit\"")
		return true
	}else if httpResponse.StatusCode == 200{	
		return false
	
} else if httpResponse.StatusCode == 498 {
	fmt.Println("\nLogin session expired!\t Please login again!\n")
	return true

} else{

	fmt.Println("something Went wrong")
	return true

}

}
