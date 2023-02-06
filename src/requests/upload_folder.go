package requests

import (
    "bytes"
    "fmt"
    "io"
    "mime/multipart"
    "net/http"
    "os"
    "path/filepath"
    "io/ioutil"
    sh "github.com/VaradBelwalkar/go_client/session_handling"
)


func UploadFolder(folderPath string) {
    data := make(map[string]interface{})

    files, err := ioutil.ReadDir(folderPath)
    if err != nil {
        panic(err)
    }

    for _, file := range files {
        if file.IsDir() {
            UploadFolder(filepath.Join(folderPath, file.Name()))
        } else {
            filePath := filepath.Join(folderPath, file.Name())
            file, err := os.Open(filePath)
            if err != nil {
                panic(err)
            }
            defer file.Close()

            buffer := new(bytes.Buffer)
            _, err = io.Copy(buffer, file)
            if err != nil {
                panic(err)
            }

            data[filePath] = buffer
        }
    }

    status, check := sh.POST_Request("/upload_folder", data)
    if check != 200 {
        if check == 500 {
            fmt.Println("Server error!")
            return
        } else if check == 502 {
            fmt.Println("Server not reachable!")
            return
        }
    }

    fmt.Printf("Folder %s uploaded successfully with status %d\n", folderPath, status)
}
