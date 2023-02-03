package session_handling



import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
	//"github.com/VaradBelwalkar/go_client/main"
)


//JWT will be in the following format

//jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"






//The JWT standard defines a specific structure for the encoded token, which consists of three parts separated by dots:

//The first part is the header, which contains information about how the JWT is encoded.
//The second part is the payload, which contains the claims. 
//Claims are statements about an entity (typically, the user) and additional data. Claims are encoded as a JSON object.
//The third part is the signature, which is used to verify that the sender of the JWT is who it says
//it is and to ensure that the message wasn't changed along the way.


//To encode a JWT, you typically need to do the following:

//Create the header and payload as JSON objects.
//Base64-encode the header and payload.
//Concatenate the encoded header, payload, and a secret key with dots (.) to form the JWT.



// Standerdized and returns type of map[string]interface{}
// To be used after successful login and JWT retrieval
func GET_Request(request_path string) map[string]interface{} {

	_, ok := os.LookupEnv("JWT")
	if ok==false{
		fmt.Println("Authentication Error!\n Please Login again")
		return nil
	}
	JWT:=os.Getenv("JWT")

	credHolder,err:=Show_Credentials()
	if err!=nil{
		fmt.Println(err)
		return nil
	}

	req, err := http.NewRequest("GET", credHolder["url"]+":"+credHolder["port"]+request_path,nil)
	client:=&http.Client{}
	req.Header.Set("Authorization","Bearer "+JWT) // JWT must be available

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	status,str:=Handle_resp_err(res)
	if status==401{
		check,str:=Login()
		if check!=false{
			fmt.Println(str)
			return nil
		}
	}else{
		fmt.Println(str)
		return nil
	}

	// Read the response body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Unmarshal the response body into a map interface 
	var response map[string]interface{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return response

}




// To be used after successful login and JWT retrieval
func POST_Request(request_path string, data map[string]interface{}) map[string]interface{} {
	b, err := json.Marshal(data)
	client:=&http.Client{}
	if err != nil {
		fmt.Println("something went wrong")
	}
	//Change URL here

	credHolder,err:=Show_Credentials()
	if err!=nil{
		fmt.Println(err)
		return nil
	}

	req, err := http.NewRequest("POST",  credHolder["url"]+":"+credHolder["port"]+request_path, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Add the JWT to the request header
		_, ok := os.LookupEnv("JWT")
	if ok==false{
		fmt.Println("Authentication Error!\n Please Login again")
		return nil
	}
	JWT:=os.Getenv("JWT")
	req.Header.Set("Authorization", "Bearer "+JWT)

	// Add the cookie to the request
	//req.AddCookie(&http.Cookie{Name: "session_id", Value: "123456"})

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	status,str:=Handle_resp_err(res)
	if status ==200{}else if status==401{
		check,str:=Login()
		if check!=false{
			fmt.Println(str)
			return nil
		}
	}else{
		fmt.Println(str)
		return nil
	}
	// Read the response
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Unmarshal the response into a Response struct
	var response map[string]interface{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return response

	
}