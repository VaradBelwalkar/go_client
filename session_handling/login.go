package session_handling

import (
	"encoding/json"
	"fmt"
	"os"
	"net/url"
	"io/ioutil"
	"strings"
	"net/http"
	"time"
	"github.com/PuerkitoBio/goquery"
	//"github.com/VaradBelwalkar/go_client"
)

		//Content-Type is one of the headers available in http request

//	Content-Type -------> application/json  

//To pass normal json data 


//Content-Type --------> application/x-www-form-urlencoded

//To recogise on server side that the request is holding form data (data submitted through form)

//Here the data is submitted as a string(obviously)as it is NOT mentioned like this "application/json"
//In a URL-encoded query string, the characters in the string are encoded so that they can be safely transmitted over the Internet. 
//The encoding replaces certain characters with a percent sign followed by a two-digit hexadecimal number. 
//or example, a space is encoded as "%20" and a plus sign is encoded as "%2B".

//For example, consider an HTML form with two text fields, 
//"name" and "email", and a submit button. 
//When the user fills out the form and clicks the submit button, the form data is sent to the server in the body of a POST request. 
//The "Content-Type" header of the request would be set to "application/x-www-form-urlencoded" and the body of the 
//request would look something like this:

//name=John+Doe&email=johndoe%40example.com

//This format is useful because it is simple and easy to parse on the server side.
//However, it has a limitation in that it can only transmit ASCII characters and does not support file uploads. 
//For more advanced functionality, other media types such as "multipart/form-data" or "application/json" may be used instead.









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








									//Understand how CSRF works


//First we make get request to the /login page to the server, then server sends the form with the CSRF token embedded into it
//You don't need to retrieve that CSRF token any way, as it is in the form itself and when you submit it ,
//with the POST request, the CSRF token gets sent with the Form data 
//You just need to retrieve the Form from html and fill in appropriate values in it and then just submit it 

//It is important to make get request and get the form as ,

//Directly sending POST request without CSRF token will not work

//CSRF tokens can be passed as header from the server but,
//They are mostly embedded in the html documents within form field like

// 					<form action="/login" method="post">
// 					  <!-- Other form fields go here -->
// 					  <input type="hidden" name="csrf_token" value="abc123">
// 					  <button type="submit">Log In</button>
// 					</form>


//It is common for servers to send the JSON Web Token (JWT) to the client in the response body of the login request. 
//The client can then extract the JWT from the response and store it locally for use in authenticating subsequent requests.

//For example, the server may return the JWT in the response body as a JSON object:

//   {
//  	"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
//   }
  
//The client can then parse the JSON response and extract the "access_token" field to obtain the JWT.

//Alternatively, the server may send the JWT in the "Authorization" header of the response. For example:

// Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c



//We will assume server is going to send the JWT token through header 



//stores the credentials, password as hash
//stores username,password,url and port of the server 
var user_credentials map[string]string

func read_credentials(){
	// Open the file in binary mode
	file, err := os.Open("credentials.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the file into a byte slice
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse the JSON data	
	err = json.Unmarshal(data, &user_credentials)
	if err != nil {
		fmt.Println(err)
		return
	}

}

// This function logs into the server and preserves JWT for further communication
func Login() (bool,string){
	// Create a new HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
 
	//Do IO on file to retrieve username and password 

	user_credentials,err:=Show_Credentials()
	if err!=nil{
		//handle error
	}
	//Do whenever submitting form data
	data := url.Values{}

	

	//Request made to get the form required
	resp,err:=http.Get(user_credentials["url"]+":"+user_credentials["port"]+"/login")
	

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("something went wrong")
		os.Exit(0)
	}

	// Find the hidden field with the name "csrf_token"
	csrfToken := doc.Find("input[name=csrf_token]").First().AttrOr("value", "")
	if csrfToken == "" {
		fmt.Println("CSRF token not found")
		return true,"\nserver side error!\n" // server side error!
	}

	//Preparing the body of the POST request, which is nothing but form data being sent using appropriate header
	data.Add("username", string(user_credentials["username"]))
	data.Add("password", string(user_credentials["password"])) //To be retrieved 
	data.Add("csrf_token",csrfToken)

	req,err:= http.NewRequest("POST",user_credentials["url"]+":"+user_credentials["port"]+"/login",strings.NewReader(data.Encode()))
	if err!=nil{
		return true,"\nServer not responding!\n" //server not responding !
	}
	//The header is set to this to recognise that the body of the request is holding form data
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	
	//Here the request is being actually sent
	//the response object will contain the JWT token
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return true,"\nSomething went wrong!\n"
	}
	defer res.Body.Close()
	
	//We can get here statuses only from 403 or 208 
	info:=Handle_resp_err(res)
	if info!=false{
		return true,""
	}
	//The JWT token
	JWT:= res.Header.Get("authorization")    //Here you can access this token anywhere in this package
	splitToken:=strings.Split(JWT, "Bearer ")
	tokenString:=splitToken[1]
	os.Setenv("JWT",tokenString)

		


//Login completed

return true,""

}






