all: mycd test

cd:
	@echo "building cd"
	go build main.go 

test:
	@echo "building unit test"
	go test -v ./mycd

clean:
	rm cd test