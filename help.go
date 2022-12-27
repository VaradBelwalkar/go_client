package github.com/VaradBelwalkar/go_client
 
//Just used to export information
var help = `
NAME
       go : Client to access microservices

DESCRIPTION
go :
         go is a way to access secure personal storage at the server with ability to access various runtimes 
         the server provides with access to some services

go STARTUP 
       To get into go environment, just type,

       >go start

go COMMANDS

       go > config

               Displays the configured URL, username, and password of the client

       go > config edit

              Edits configuration and saves the changes
	   
	go > version

               Shows the version number of the project

       go > server

               Shows information about the server

       go > signup

              Creates an account with given username and password

       go > view
      
              lists all the files in the cloud

       go > uploadfile <file-path>

              uploads the specified file into the cloud

       go > uploaddir <dir-path>

              uploads the specified directory to the cloud

       go > download <filename> <filepath>

              downloads the specified file from the cloud

       go > delete <filepath>+'/'+<filename>

              deletes the specified file in the cloud

       go > sync <dirpath>

              sync the specified directory with the cloud



go COMMANDS TO ACCESS RUNTIMES

  (CURRENTLY AVAILABLE RUNTIMES : ubuntu , development_server)

  (You can own maximum of 5 containers at a time)


       go > container run <runtime_name> 
             
              get a new specifed container

       go > container list images

              get available os images

       go > container list containers

              get the list of containers that you own (useful for which container to resume)

       go > container resume <container_name>

              get the specified container that you own from your owned_containers list

       go > container stop <container_name>

              stop the specified container or all  on server side

       go > container remove <container_name>

              remove the specified container or all from the server

`


func Help(){
       fmt.Println(help)
}
