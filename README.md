# HttpsClient

This repository contains software to test client certificates against a web server with a configured certificate authority.
The assumption is that you have created an RSA private key (client.key) file and you have had a certificate signing request signed by a certificate authority making an x509 certificate (client.crt) file.

## Build

This program can be built using the makefile provided.  There are commands for building OSX, Linux and Windows versions of the software.  Binaries are generated in the bin directory.  This software is built with Go 1.5.  
Alternatively you can download the binaries you need from the bin directory.

    // OSX
    make osx
    
    // Linux
    make linux
    
    // Windows 32 bit
    make win32
    
    // Windows 64 bit
    make win64

## Usage

The httpsClient program requires three arguments:

    - key  The key file created by the client.
    - cert The certificate generated by the certificate authority.
    - url  The url of the web server to test the client certificate against. 
 
## Example

    ./httpsClient_linux -key client.key -cert client.crt -url https://www.someserver.com
    
    // Sample output
    httpsClient v1.0 - UK Home Office
    
    STATUS
    HTTP/1.1 200 OK
    
    HEADERS
    Server: nginx
    Date: Wed, 13 Jan 2016 10:02:52 GMT
    Content-Type: text/plain;charset=ISO-8859-1
    Content-Length: 7
    Connection: keep-alive
    X-Application-Context: application
    
    BODY
    SUCCESS

    

