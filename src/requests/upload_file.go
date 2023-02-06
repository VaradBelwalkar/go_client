package requests

import (
    "bytes"
    "io"
    "net/http"
    "os"
    "fmt"
    sh "github.com/VaradBelwalkar/go_client/session_handling"
)



func UploadFile(filepath string) {
    // Open the file that you want to send in the request
    file, err := os.Open(filepath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Create a new buffer to hold the contents of the file
    buffer := new(bytes.Buffer)

    // Copy the contents of the file into the buffer
    _, err = io.Copy(buffer, file)
    if err != nil {
        panic(err)
    }

    var data map[string]interface{}
    data["file"]=buffer
    // Create a new HTTP POST request
    _,check:=sh.POST_Request("/upload_file",data)
    if check!=200{
        if check==500{fmt.Println("Server error!")
        return
    } else if check == 502{
        fmt.Println("Server not reachable!")
        return
    }
        
    }
    fmt.Println("File uploaded Successfully!")

  
}

