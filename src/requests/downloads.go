package requests

import (
	"log"
	"fmt"
	"os/exec"
	"os"
	"strings"
	sh "github.com/VaradBelwalkar/go_client/session_handling"
)
func Downloads(fileOrFolder string,containerPath string,userPath,containerName string){
	colorReset := "\033[0m"
    colorYellow := "\033[33m"
	user_credentials,err:=sh.Show_Credentials()
	if err!=nil{
		fmt.Println(string(colorYellow),"Please run change config to store your credentials",string(colorReset))
	}
	scriptPath := sh.ProjectPath+"/connections/download_script.sh"
	parts := strings.Split(containerName, "_")
	port := parts[1]
	cmd := exec.Command(scriptPath,fileOrFolder,containerPath,userPath,port,user_credentials["ip"],sh.ProjectPath+"/connections/keyForRemoteServer")
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



// 	download <file,folder> <path in container> <path in your computer> <containername> 