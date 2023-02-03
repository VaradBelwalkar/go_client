package session_handling

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
	"strings"
	"net/url"
	"time"
	"github.com/PuerkitoBio/goquery"
)




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
        fmt.Println("Your Password does not match. Please try again")
		fmt.Print("Enter your password: ")
		password, _ = reader.ReadString('\n')
		fmt.Print("confirm your password: ")
		password1, _ = reader.ReadString('\n')
	}

	fmt.Print("Enter the server IP: ")
	url,_:=reader.ReadString('\n')
	fmt.Print("Enter the port: ")
	port,_:=reader.ReadString('\n')
	//Request made to get the form required	

	urlString:=	"http://"+strings.ReplaceAll(url, " ", "")+strings.ReplaceAll(port, " ", "")+"/register"
	res,err:=http.Get(urlString)
	

	doc, err := goquery.NewDocumentFromReader(res.Body)
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

	user_credentials,err:=Show_Credentials()
	if err!=nil{
		//handle error
	}

	req,err:= http.NewRequest("POST",user_credentials["url"]+":"+user_credentials["port"]+"/register",strings.NewReader(data.Encode()))
	if err!=nil{
		fmt.Println(err)
		return 
	}
	//The header is set to this to recognise that the body of the request is holding form data
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	
	//Here the request is being actually sent
	//the response object will contain the JWT token
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	
	// Unmarshal the response into a Response struct
	var response http.Response
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	//The JWT token
	 //Here you can access this token anywhere in this package
    

	//Here the 201 StatusCode means the resource is successfully created on the server
	if response.StatusCode == 201 {
		//Meaning the registration is successful
		fmt.Println("The registered successfully!")
		Store_credentials(username,password,url,port)
		return

	} else if response.StatusCode == 409 {  //    409 StatusCode indicates a "Conflit" that server cannot create a resource because
											  //    it already exists
		fmt.Println("The username already exists! Please choose another username")
		goto try

	} else {
		fmt.Println(response.StatusCode)
	}






}















