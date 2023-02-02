package session_handling

import (
	"encoding/json"
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
	//"github.com/VaradBelwalkar/go_client/main"
)


func Store_credentials(username string,password string,url string,port string) {
	user_credentials:= map[string]interface{}{
		"usernaewme":username,
		"password":password,
		"url":url,
		"port":port,
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
		return
	}
	fmt.Println("Configuration stored successfully!")


}



func Setup(){
    reader := bufio.NewReader(os.Stdin)
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
        fmt.Println("Your Password does not match. Please try again")
		fmt.Print("Enter your password: ")
		password, _ = reader.ReadString('\n')
		fmt.Print("confirm your password: ")
		password1, _ = reader.ReadString('\n')
	}

	fmt.Print("Enter the url: ")
	url,_:=reader.ReadString('\n')
	fmt.Print("Enter the port: ")
	port,_:=reader.ReadString('\n')


	Store_credentials(username,password,url,port)


}


func Show_Credentials()(map[string]string,error){
	var user_credentials map[string]string
	// Open the file in binary mode
	file, err := os.Open("credentials.bin")
	if err != nil {
		fmt.Println("File not found!\n \t\t Run `change config` to configure user credentials")
		return nil,err
	}
	defer file.Close()

	// Read the file into a byte slice
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil,err
	}

	// Parse the JSON data	
	err = json.Unmarshal(data, &user_credentials)
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	//has username, password URL, and port
	return user_credentials,nil

}