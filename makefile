
init:
	dep ensure -v -update
run:
	port=3001 go run main.go 
build: 
	make init
	go build -v -o dist/whoisServer