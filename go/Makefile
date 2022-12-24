all: mycd test

mycd:
	@echo "building cd"
	go build main.go
	cp main ../bin/mycd

test:
	@echo "building unit test"
	go test -v github.com/monu1122/mycd/go/mycd

clean:
	rm -f main ../bin/mycd