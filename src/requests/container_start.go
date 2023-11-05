package requests

import (
	"fmt"
	"log"
	"os/exec"
	"bufio"
	"os"
	"io/ioutil"
	sh "github.com/VaradBelwalkar/go_client/session_handling"
)


func Container_Start(imageName string,browser bool,given_port string){
	reader := bufio.NewReader(os.Stdin)
    colorReset := "\033[0m"
	colorGreen := "\033[32m"
    colorRed := "\033[31m"
	colorBlue := "\033[34m"
    colorYellow := "\033[33m"
	request_path:="/container/resume/"+imageName


	//resp is of type map[string]interface{}
	resp,status:= sh.GET_Request(request_path)  

	if status!=200 {
		  if status == 500{
		fmt.Println(string(colorRed),"Server error!",string(colorReset))
		return
		} else if status == 502{
			return
		}else if status == 401{
			fmt.Println(string(colorRed),"Something went wrong on your side!",string(colorReset))
			return
		} else if status == 404{
			fmt.Println(string(colorYellow),"No such container!",string(colorReset))
			return
		}
	}
	
	privateKey:=resp["privatekey"].(string)	
	container_ip:=resp["container_ip"].(string)
	 //define the path to the bash script
	
	err := ioutil.WriteFile(sh.ProjectPath+"/keyForRemoteServer", []byte(privateKey), 0600)
    if err != nil {
        fmt.Println(string(colorRed),"Something went wrong while storing PrivateKey",string(colorReset))
		return
    }
	// Parameters to pass to the script
	fmt.Println(string(colorBlue),"Copy the following line in your VSCode to get development environment:\n",string(colorReset))
	fmt.Println(string(colorGreen),"ssh "+"-i "+sh.ProjectPath+"/keyForRemoteServer "+" root@"+container_ip,string(colorReset))
	fmt.Println(string(colorBlue),"\nPress Enter when done:",string(colorReset))
	_,_=reader.ReadString('\n')
	// start the script
	cmd := exec.Command("ssh","-i",sh.ProjectPath+"/keyForRemoteServer","root@"+container_ip)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if browser == true{
		browser_cmd := exec.Command("google-chrome", "--new-tab",container_ip+":"+given_port)
		browser_cmd.Stdin = os.Stdin
		browser_cmd.Stdout = os.Stdout
		browser_cmd.Stderr = os.Stderr
		
		// start the script and wait for it to finish
		if err := browser_cmd.Start(); err != nil {
			// handle error
			log.Fatal(err)
		}
		if err := browser_cmd.Wait(); err != nil {
			// handle error
			log.Fatal(err)
		}
	}
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

