# clean the cache
clean:
	-go clean -cache 

# builds the binary
build:
	go build

# checks code quality 
check:
	go vet -asmdecl -bools -assign ./

# to run all the tests in the project
test:
	cd handlers/tests && ginkgo -v
	cd db/tests && ginkgo -v

reload:
	sudo docker-compose build 
	docker-compose up

start:
	docker-compose up