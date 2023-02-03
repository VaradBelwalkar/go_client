package session_handling

import(
	"bufio"
	"fmt"
	"os"
	"net/http"
	"strings"
	"net/url"
	"time"
	"github.com/PuerkitoBio/goquery"
)



func Remove_account() {
	// Create a new HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
 	
	//Do IO on file to retrieve username and password 

	//Do whenever submitting form data
	data := url.Values{}

	
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the username : ")
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
	user_credentials,err :=Show_Credentials()
	var url,port string
	if err!=nil{
		fmt.Print("Enter the server IP: ")
		url,_=reader.ReadString('\n')
		fmt.Print("Enter the port: ")
		port,_=reader.ReadString('\n')
	}else{
		url=user_credentials["url"]
		port=user_credentials["port"]
	}


	//Request made to get the form required	

	urlString:=	"http://"+strings.ReplaceAll(url, " ", "")+strings.ReplaceAll(port, " ", "")+"/remove_account"
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

	req,err:= http.NewRequest("POST",user_credentials["url"]+":"+user_credentials["port"]+"/remove_account",strings.NewReader(data.Encode()))
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
	
	status,str:=Handle_resp_err(res)
	if status == 200{
		fmt.Println("Account deleted successfully!")
		return
	}else{
		if status == 403{
			fmt.Println("\nWrong password!")
			return
		} else if status == 404{
		fmt.Println("\nUser doesn't exist!")}
		fmt.Println(str)
		return
	}


}



