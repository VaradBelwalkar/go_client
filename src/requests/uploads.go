package requests

import (
	"os/exec"
	"os"
	"strings"
	"fmt"
	sh "github.com/VaradBelwalkar/go_client/session_handling"
)
func Uploads(fileOrFolder string,localPath string,containerPath,containerName string){
	colorReset := "\033[0m"
    //colorYellow := "\033[33m"
	colorRed := "\033[31m"
	
	parts := strings.Split(containerName, "_")
	container_ip := parts[1]
	cmd := exec.Command("scp","-i",sh.ProjectPath+"/keyForRemoteServer",localPath,"root@"+container_ip+":"+containerPath)
	if fileOrFolder == "file"{
	
	}else if fileOrFolder == "folder"{
		cmd = exec.Command("scp","-r","-i",sh.ProjectPath+"/keyForRemoteServer",localPath,"root@"+container_ip+":"+containerPath)
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

// upload <file,folder> <path> <containername>