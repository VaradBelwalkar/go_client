package requests

import (
	"fmt"
	"log"
	"os/exec"
	"os"
	"io/ioutil"
	sh "github.com/VaradBelwalkar/go_client/session_handling"
)


func Container_Run(imageName string){

	request_path:="/container/run/"+imageName


	//resp is of type map[string]interface{}
	resp,status:= sh.GET_Request(request_path)  

	if status!=200 {
		fmt.Println("something went wrong!")
		return 
	}

	privateKey:=resp["privatekey"].(string)	
	port:=resp["port"].(string)
	// define the path to the bash script
	scriptPath := "/home/varad/repositories/go_client/src/connections/bash_script.sh"
	
	err := ioutil.WriteFile("/home/varad/repositories/go_client/src/connections/keyForRemoteServer", []byte(privateKey), 0644)
    if err != nil {
        panic(err)
    }
	// Parameters to pass to the script
	params := []string{port}
	
	// start the script
	cmd := exec.Command(scriptPath, params...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	// start the script and wait for it to finish
	if err := cmd.Start(); err != nil {
		// handle error
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		// handle error
		log.Fatal(err)
	}
	
	
}

