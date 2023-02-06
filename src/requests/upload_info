//multipart/form-data is a content type used to submit data as part of a form request. 
//It is often used when submitting HTML forms that include file inputs, 
//because it allows the form data and the files to be sent in the same request.

//When a form is submitted using multipart/form-data, the form data is included in the body of the request as a series of "parts". 
//Each part can contain either form data or the contents of a file. The parts are separated by a boundary string, 
//which is specified in the Content-Type header of the request.

//Here's an example of a multipart/form-data POST request that includes both form data and a file:

//Copy code
//POST /upload HTTP/1.1
//Content-Type: multipart/form-data; boundary=---------------------------7da24f2e50046
//Content-Length: 554

//-----------------------------7da24f2e50046
//Content-Disposition: form-data; name="field1"

//value1
//-----------------------------7da24f2e50046
//Content-Disposition: form-data; name="field2"

//value2
//-----------------------------7da24f2e50046
//Content-Disposition: form-data; name="file1"; filename="file.txt"
//Content-Type: text/plain

//This is the contents of the file.

//-----------------------------7da24f2e50046--
//In this example, the request includes three parts: two form data parts (field1 and field2) and a file part (file1).
// The boundary string is ---------------------------7da24f2e50046.

//The mime/multipart package in Go provides functions for creating and parsing multipart/form-data requests and responses.