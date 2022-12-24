all: mycd test

mycd:
	@echo "building mycd"
	g++ main.cpp -o mycd

test:
	@echo "building unit test"
	g++ mycd_ut.cpp -o test

clean:
	rm mycd test