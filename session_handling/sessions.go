package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

		//Content-Type is one of the headers available in http request

//	Content-Type -------> application/json  

//To pass normal json data 


//Content-Type --------> application/x-www-form-form-urlencoded

//To recogise on server side that the request is holding form data (data submitted through form)




		// Authorization is one of the headers available in the http request
	
//  Authorization -----------> "Bearer <JWT-Token>"

//This indicates the authorization method, which is JWT one



	//Bearer is used indicate JWT based authentication while Basic is used to indicate username:password based authentication

//Basic authentication: Basic authentication is a simple authentication scheme that uses a username and password to authenticate requests. 
// To use basic authentication, you can set the value of the Authorization header to "Basic <credentials>",
// where <credentials> is the base64-encoded string of the username and password separated by a colon. For example:
// Authorization: Basic QWxhZGRpbjpvc


				//CSRF Token 

//Cross-Site Request Forgery (CSRF) tokens are usually passed in a header or as a request parameter in the query string.
//The specific location of the CSRF token will depend on the requirements of the server and the client.
//Here we are passing CSRF Token as header in the request field






// Request represents a request to the server
type Request struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Response represents a response from the server
type Response struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func main() {
	// Create a new HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Set up the request
	reqBody, err := json.Marshal(Request{ID: 1, Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}
	

	//Do IO on file to retrieve username and password 



	//Do whenever submitting form data
	data := url.Values{}
	data.Add("username", "John Doe")
	data.Add("passwordhash", "lskadfjlkasdht3erw") //To be retrieved 

	req,err:= http.NewRequest("POST","http://url/login",string.NewReader(data.Encode()))
	if err!=nil{
		fmt.Println(err)
		return 
	}
	//The header is set to this to recognise that the body of the request is holding form data
	req.Header.Set("Content-Type","application/x-www-form-form-urlencoded")
	
	//Here the request is being passed 
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	
	//Once logged-in, store the JWT token
	//
	//
	//




	req, err := http.NewRequest("POST", "http://example.com/login", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Add the JWT to the request header
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	req.Header.Set("Authorization", "Bearer "+jwt)

	// Add the cookie to the request
	req.AddCookie(&http.Cookie{Name: "session_id", Value: "123456"})

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// Read the response
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the response into a Response struct
	var response Response
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}
