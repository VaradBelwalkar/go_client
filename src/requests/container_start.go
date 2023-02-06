package requests

import (
	"fmt"
	"io/ioutil"
	"log"
	"bufio"
	"os"
	"os/exec"
	"golang.org/x/crypto/ssh"
	sh "github.com/VaradBelwalkar/go_client/session_handling"
)


func Container_Start(containerName string){

	request_path:="resume/"+containerName


	//resp is of type map[string]interface{}
	resp,err:= sh.GET_Request(request_path)
	if err!=nil{
		fmt.Println("some error here to be handled here")
		return 
	}
	privateKey:=resp["privatekey"]	
	port:=resp["port"]
	// define the path to the bash script
	scriptPath := "/path/to/script.sh"

	// Parameters to pass to the script
	params := []string{privateKey,port}
	
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

