all:
	env GO111MODULE=on go install -v .

.PHONY: clean test update-modules
clean:
	rm -f bitumwallet

test:
	env GO111MODULE=on ./run_tests.sh

update-modules:
	env GO111MODULE=on go get -u
	env GO111MODULE=on go mod tidy -v
