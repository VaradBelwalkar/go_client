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


func store_credentials(username string,password string) {
	user_credentials:= map[string]interface{
		"username":username,
		"password":password
	}

	//This json Marshalling creates an array of unit8
	// That is info is of []unit8 type    Here []uint8 is same as []byte 
	info_in_bytes,err :=json.Marshal(user_credentials)
	if err!=nil{
		panic(err)
	}

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
	fmt.Println("Configuration stored successfully!")


}



func Setup(){

	fmt.Print("Enter the username: ")
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
	store_credentials(username,password)


}


func Show_Credentials(){
	var user_credentials map[string]interface{}
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

	fmt.Println("username : ",user_credentials["username"])
	fmtl.Println("password :",user_credentials["password"])

}