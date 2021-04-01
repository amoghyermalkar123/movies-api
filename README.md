# movies-api
Task for building Movies API (like IMDB)

# Setting up the Project :
- You will need a local postgres instance up and running with following user-credentials:
  -  ##### username : amogh 
  -  ##### password : amogh 

- Once your local postgres database is running, you can setup the project.
  - To do that `cd` into the root of the project and run 
    `make setup`
  - This will download all the dependencies that are required to run the project.

- To start up the HTTP server run `make start`

# Some additional commands to go over for the project :
- `make build`
  - builds the project binary
- `make test`
  - runs tests
- `make check`
  - runs a code sanity check script
