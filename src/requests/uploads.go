package requests

import (
	"log"
	"os/exec"
	"os"
	"strings"
	"fmt"
	sh "github.com/VaradBelwalkar/go_client/session_handling"
)
func Uploads(fileOrFolder string,localPath string,containerPath,containerName string){
	colorReset := "\033[0m"
    colorYellow := "\033[33m"
	user_credentials,err:=sh.Show_Credentials()
	if err!=nil{
		fmt.Println(string(colorYellow),"Please run change config to store your credentials",string(colorReset))
	}
	scriptPath := sh.ProjectPath+"/connections/upload_script.sh"
	parts := strings.Split(containerName, "_")
	port := parts[1]
	cmd := exec.Command(scriptPath,fileOrFolder,localPath,containerPath,port,user_credentials["ip"],sh.ProjectPath+"/connections/keyForRemoteServer")
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

// upload <file,folder> <path> <containername>