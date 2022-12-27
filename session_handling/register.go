package session_handling

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"golang.org/x/net/html"
	"github.com/PuerkitoBio/goquery"
	"github.com/VaradBelwalkar/go_client/main"
)




//stores the credentials, password as hash


func store_credentials(username string,password string) {
	user_credentials:= map[string]interface{
		"username":username,
		"password":password
	}

	//This json Marshalling creates an array of unit8
	// That is info is of []unit8 type    Here []uint8 is same as []byte 
	info_in_bytes,err :=json.Marshal(user_credentials)


	f, err := os.OpenFile("credentials.bin", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		fmt.Println("Something went wrong while storing credentials!Please try again")
	}
	defer f.Close()
	
	_, err = f.Write(info_in_bytes)
	if err != nil {
		panic(err)
		return
	}



}




// This function logs into the server and preserves JWT for further communication
func Register() {
	// Create a new HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
 	
	//Do IO on file to retrieve username and password 

	//Do whenever submitting form data
	data := url.Values{}

	
	reader := bufio.NewReader(os.Stdin)
try:
	fmt.Print("Enter the username you want: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter your password: ")
	password, _ := reader.ReadString('\n')
	fmt.Print("confirm your password: ")
	password1, _ := reader.ReadString('\n')

    for {
		if password == password1{
			break;
		}
        fmt.Println("Your Password does not match. Please try again\n")
		fmt.Print("Enter your password: ")
		password, _ := reader.ReadString('\n')
		fmt.Print("confirm your password: ")
		password1, _ := reader.ReadString('\n')
	}
	
	//Request made to get the form required
	resp,err:=http.NewRequest("http://url/Register")
	

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find the hidden field with the name "csrf_token"
	csrfToken := doc.Find("input[name=csrf_token]").First().AttrOr("value", "")
	if csrfToken == "" {
		fmt.Println("CSRF token not found")
		return
	}

	//Preparing the body of the POST request, which is nothing but form data being sent using appropriate header
	data.Add("username", username)
	data.Add("password", password) 
	data.Add("csrf_token",csrfToken)

	req,err:= http.NewRequest("POST","http://url/login",string.NewReader(data.Encode()))
	if err!=nil{
		fmt.Println(err)
		return 
	}
	//The header is set to this to recognise that the body of the request is holding form data
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	
	//Here the request is being actually sent
	//the response object will contain the JWT token
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	
	// Unmarshal the response into a Response struct
	var response Response
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		panic(err)
	}

	//The JWT token
	 //Here you can access this token anywhere in this package
    

	//Here the 201 StatusCode means the resource is successfully created on the server
	if response.StatusCode == 201 {
		//Meaning the registration is successful
		fmt.Println("The registered successfully!")
		store_credentials(username,password)
		return

	} else if response.StatusCode == 409 {  //    409 StatusCode indicates a "Conflit" that server cannot create a resource because
											  //    it already exists
		fmt.Println("The username already exists! Please choose another username\n")
		goto try

	} else {
		fmt.Prinln(response.StatusCode)
	}






}









