package main

import (
    "bytes"
    "io"
    "net/http"
    "os"
)

func main() {
    // Open the file that you want to send in the request
    file, err := os.Open("path/to/file.txt")
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

    // Create a new HTTP POST request
    req, err := http.NewRequest("POST", "http://localhost:8080/upload", buffer)
    if err != nil {
        panic(err)
    }

    // Set the Content-Type header to "text/plain"
    req.Header.Set("Content-Type", "text/plain")

    // Send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}
