package session_handling

import (
	"encoding/json"
	"fmt"
	"bufio"
	"strings"
	"os"
	"io/ioutil"
	//"github.com/VaradBelwalkar/go_client/main"
)


func Store_credentials(username string,password string,IP string,port string) {
	user_credentials:= map[string]string{
		"username":username,
		"password":password,
		"ip":strings.ReplaceAll(IP," ",""),
		"port":port,
	}

	//This json Marshalling creates an array of unit8
	// That is info is of []unit8 type Here []uint8 is same as []byte 
	info_in_bytes,err :=json.Marshal(user_credentials)
	if err!=nil{
		panic(err)
	}

	f, err := os.OpenFile("/home/varad/repositories/go_client/credentials.bin", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
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
	var tempUsername string
	var username string
	var tempPassword string
	var password string
	var tempPassword1 string
	var password1 string
	var tempIP string
	var IP string 
	var tempPort string
	var port string
	var _ error
	for {
	fmt.Print("Enter the username: ")
	tempUsername, _ = reader.ReadString('\n')
	username=strings.ReplaceAll(tempUsername,"\n","")
		if username!=""{break
		}else {
			fmt.Println("Please enter a valid username!")
		}

	}
	for{
	fmt.Print("Enter your password: ")
	tempPassword, _ = reader.ReadString('\n')
	password=strings.ReplaceAll(tempPassword,"\n","")
	if password!=""{break
		}else {
			fmt.Println("Please enter a valid password!")
		}
	}
	for{
	fmt.Print("confirm your password: ")
	tempPassword1, _ = reader.ReadString('\n')
	password1=strings.ReplaceAll(tempPassword1,"\n","")
	if password1!=""{break
	} else{
		fmt.Println("Please enter valid confirmation!")
	}

	}
    for {
		if password == password1{
			break;
		}
        fmt.Println("Your Password does not match. Please try again")
		for{
			fmt.Print("Enter your password: ")
			tempPassword, _ = reader.ReadString('\n')
			password=strings.ReplaceAll(tempPassword,"\n","")
			if password!=""{break
				}else {
					fmt.Println("Please enter a valid password!")
				}
			}
			for{
			fmt.Print("confirm your password: ")
			tempPassword1, _ = reader.ReadString('\n')
			password1=strings.ReplaceAll(tempPassword1,"\n","")
			if password1!=""{break
			} else{
				fmt.Println("Please enter valid confirmation!")
			}
		
			}
	}
	for{
	fmt.Print("Enter the IP: ")
	tempIP,_=reader.ReadString('\n')
	IP=strings.ReplaceAll(tempIP,"\n","")
		if IP!=""{
			break
		}	else{
			fmt.Println("Please enter valid IP")
		}
}	
for{
	fmt.Print("Enter the port: ")
	tempPort,_=reader.ReadString('\n')
	port=strings.ReplaceAll(tempPort,"\n","")
	
	if port!=""{
		break
	}else{
		fmt.Println("Please enter a valid port number")
	}
}
	Store_credentials(username,password,IP,port)


}


func Show_Credentials()(map[string]string,error){
	var user_credentials map[string]string
	// Open the file in binary mode
	file, err := os.Open("/home/varad/repositories/go_client/credentials.bin")
	if err != nil {
		fmt.Println("\t\tFile not found!\n \t\t Run `change config` to configure user credentials")
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
		return nil,err
	}
	//has username, password URL, and port
	return user_credentials,nil

}


func Set_url(){
	reader := bufio.NewReader(os.Stdin)
	var tempIP string
	var IP string
	for{
		fmt.Print("Enter the IP: ")
		tempIP,_=reader.ReadString('\n')
		IP=strings.ReplaceAll(tempIP,"\n","")
			if IP!=""{
				break
			}	else{
				fmt.Println("Please enter valid url")
			}
	}

	var user_credentials map[string]string
	// Open the file in binary mode
	file, err := os.Open("/home/varad/repositories/go_client/credentials.bin")
	if err != nil {
		fmt.Println("\t\tFile not found!\n \t\t Run `change config` to configure user credentials")
		return 
	}
	defer file.Close()

	// Read the file into a byte slice
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return 
	}

	// Parse the JSON data	
	err = json.Unmarshal(data, &user_credentials)
	if err != nil {	
		return 
	}
	Store_credentials(user_credentials["username"],user_credentials["password"],IP,user_credentials["port"])
}


func Set_port(){

	reader := bufio.NewReader(os.Stdin)
	var PORT string
	for{
		fmt.Print("Enter the port: ")
		PORT,_=reader.ReadString('\n')
		PORT=strings.TrimSuffix(PORT, "\n")
		
		if PORT!=""{
			break
		}else{
			fmt.Println("Please enter a valid port number")
		}
	}

	var user_credentials map[string]string
	// Open the file in binary mode
	file, err := os.Open("/home/varad/repositories/go_client/credentials.bin")
	if err != nil {
		fmt.Println("\t\tFile not found!\n \t\t Run `change config` to configure user credentials")
		return 
	}
	defer file.Close()

	// Read the file into a byte slice
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return 
	}

	// Parse the JSON data	
	err = json.Unmarshal(data, &user_credentials)
	if err != nil {	
		return 
	}

	Store_credentials(user_credentials["username"],user_credentials["password"],user_credentials["ip"],PORT)




}