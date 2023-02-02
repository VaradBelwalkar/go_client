package main

import (
    "bytes"
    "fmt"
    "io"
    "mime/multipart"
    "net/http"
    "os"
    "path/filepath"
)

func main() {
    // Create a new HTTP POST request
    req, err := http.NewRequest("POST", "http://localhost:8080/upload", nil)
    if err != nil {
        panic(err)
    }

    // Set the Content-Type header to "multipart/form-data"
    req.Header.Set("Content-Type", "multipart/form-data")

    // Create a new multipart/form-data writer
    writer := multipart.NewWriter(req.Body)
    defer writer.Close()

    // Walk through the directory tree and add all files to the request
    filepath.Walk("path/to/folder", func(path string, info os.FileInfo, err error) error {
        if info.IsDir() {
            return nil
        }

        // Open the file
        file, err := os.Open(path)
        if err != nil {
            return err
        }
        defer file.Close()

        // Add the file to the request
        part, err := writer.CreateFormFile(info.Name(), info.Name())
        if err != nil {
            return err
        }
        _, err = io.Copy(part, file)
        return err
    })

    // Send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}
