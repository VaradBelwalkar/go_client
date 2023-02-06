package session_handling

import (


	"net/http"

	//"github.com/VaradBelwalkar/go_client"
)



//Handles the reponses from the server with default status codes
//It will handle and print appropriate response to the user so no handling in the client code
func Handle_resp_err(httpResponse *http.Response) (int,string){

	if httpResponse.StatusCode == 403{ // wrong password

		return 403,"\nWrong password!\tPlease update your password by running config edit\n"
	}else if httpResponse.StatusCode == 208{

		return 404,"User doesn't exist!Please create account by running \"setup\"or update your credentials by running \"config edit\""
	}else if httpResponse.StatusCode == 200{	
		return 200,"OK"
	
} else if httpResponse.StatusCode == 401 { //invalid or expired JWT

	return 401,"\nLogin session expired!\t Please login again!\n"

}else if httpResponse.StatusCode == 409{
	return 401,"\nUser already exists!"
}else if httpResponse.StatusCode == 400{
	return 400,"\nSomething went wrong on your side!"
}else{

	return 500,"something Went wrong"

}

}


// StatusConflit 409