package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	keyFile  = flag.String("key", "", "A private key file.")
	certFile = flag.String("cert", "", "A certificate file.")
	url      = flag.String("url", "", "url to test.")
)

func main() {
	fmt.Println("httpsClient v1.0 - UK Home Office\n")

	// Test flags
	flag.Parse()
	if testFlags(*keyFile, *certFile, *url) == false {
		return
	}

	// Load client key and cert files
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// Do Get
	resp, err := client.Get(*url)
	if err != nil {
		fmt.Println("No connection to", *url)
		return
	}
	defer resp.Body.Close()

	// Print response
	printResponse(resp)
}

func printResponse(resp *http.Response) {
	// Status
	fmt.Println("STATUS")
	fmt.Println(resp.Proto, resp.Status)

	// Headers
	fmt.Println("\nHEADERS")
	for k, v := range resp.Header {
		fmt.Println(k+":", v[0])
	}

	// Body
	fmt.Println("\nBODY")
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func testFlags(keyFile, certFile, url string) bool {
	// Test flags have been entered
	if len(keyFile) == 0 || len(certFile) == 0 || len(url) == 0 {
		flag.Usage()
		return false
	}

	// Test files exist
	if testFileExists(keyFile) == false || testFileExists(certFile) == false {
		return false
	}

	return true
}

func testFileExists(fileName string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Failed to open file", err)
		return false
	}
	defer file.Close()
	return true
}
