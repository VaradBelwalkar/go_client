package requests

import (
	"fmt"
	"os/exec"
	"os"
	"strings"
	sh "github.com/VaradBelwalkar/go_client/session_handling"
)
func Downloads(fileOrFolder string,containerPath string,localPath,containerName string){
	colorReset := "\033[0m"
    colorYellow := "\033[33m"
	colorRed := "\033[31m"
	user_credentials,err:=sh.Show_Credentials()
	if err!=nil{
		fmt.Println(string(colorYellow),"Please run change config to store your credentials",string(colorReset))
	}

	parts := strings.Split(containerName, "_")
	port := parts[1]
	cmd := exec.Command("scp","-i",sh.ProjectPath+"/keyForRemoteServer","-P",port,"root@"+user_credentials["ip"]+":"+containerPath,localPath)
	if fileOrFolder == "file"{
	
	} else if fileOrFolder == "folder"{
		cmd = exec.Command("scp","-i",sh.ProjectPath+"/keyForRemoteServer","-P",port,"root@"+user_credentials["ip"]+":"+containerPath,localPath)
	}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	
	// start the script and wait for it to finish
	if err := cmd.Start(); err != nil {
		// handle error
		fmt.Println(string(colorRed),"Something went wrong,\n Check ip address or port if configured correctly else might be server issue!",string(colorReset))
		return
	}
	if err := cmd.Wait(); err != nil {
		// handle error
		fmt.Println(string(colorRed),"Something went wrong,\n Check ip address or port if configured correctly else might be server issue!",string(colorReset))
		return
	}

}



// 	download <file,folder> <path in container> <path in your computer> <containername> 