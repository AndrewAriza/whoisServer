run:
	dep ensure -v -update
	port=3001 go run main.go 

build: 
	dep ensure -v -update
	go build -v -o dist/whoisServer