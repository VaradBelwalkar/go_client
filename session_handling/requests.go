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



func POST_Request(){


	req, err := http.NewRequest("POST", "http://example.com/login", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Add the JWT to the request header
	
	req.Header.Set("Authorization", "Bearer "+main.JWT)

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

//
func GET_Request(){





}