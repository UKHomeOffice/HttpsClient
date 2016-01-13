all: osx linux win32 win64

osx:
	GOOS=darwin GOARCH=amd64 go build -o bin/httpsClient_osx httpsClient.go

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/httpsClient_linux httpsClient.go

win32:
	GOOS=windows GOARCH=386 go build -o bin/httpsClient_win32.exe httpsClient.go

win64:
	GOOS=windows GOARCH=amd64 go build -o bin/httpsClient_win64.exe httpsClient.go

clean:
	rm -f bin/*


