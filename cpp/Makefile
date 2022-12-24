all: mycd test

mycd:
	@echo "building mycd"
	g++ main.cpp -o mycd
	cp mycd ../bin/mycd

test:
	@echo "building unit test"
	g++ mycd_ut.cpp -o test

clean:
	rm -f mycd test ../bin/mycd