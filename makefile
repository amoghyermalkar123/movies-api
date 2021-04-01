# clean the cache
clean:
	-go clean -cache 

# builds the binary
build:
	go build

# install all dependencies
setup:
	go mod download

# start the HTTP server
start:
	go run server.go

# checks code quality 
check:
	go vet -asmdecl -bools -assign ./

# to run all the tests in the project
test:
	cd handlers/tests && ginkgo -v
	cd db/tests && ginkgo -v

# restore database dump
restore:
	pg_restore -U postgres -d movieserver -1 dump.sql
