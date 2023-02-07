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
	var tempIP,IP,tempPort,port string
	if err!=nil{
		fmt.Print("Enter the server IP: ")
		tempIP,_=reader.ReadString('\n')
		IP=strings.ReplaceAll(tempIP,"\n","")
		fmt.Print("Enter the port: ")
		tempPort,_=reader.ReadString('\n')
		port=strings.ReplaceAll(tempPort,"\n","")

	}else{
		IP=user_credentials["ip"]
		port=user_credentials["port"]
	}


	//Request made to get the form required	

	urlString:=	"http://"+strings.ReplaceAll(IP, " ", "")+strings.ReplaceAll(port, " ", "")+"/remove_account"
	res,err:=http.Get(urlString)
	

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find the hidden field with the name "csrf_token"
	csrfToken := doc.Find("input[name=csrf]").First().AttrOr("value", "")
	if csrfToken == "" {
		fmt.Println("CSRF token not found")
		return
	}

	//Preparing the body of the POST request, which is nothing but form data being sent using appropriate header
	data.Add("username", username)
	data.Add("password", password) 
	data.Add("csrf",csrfToken)
	cookie := &http.Cookie{
        Name:   "csrftoken",
        Value:  csrfToken,
        MaxAge: 300,
    }
	req,err:= http.NewRequest("POST",urlString,strings.NewReader(data.Encode()))
	if err!=nil{
		fmt.Println(err)
		return 
	}
	req.AddCookie(cookie)
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



