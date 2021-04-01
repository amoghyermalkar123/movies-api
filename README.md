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

# Folder structure : (brief overview)
- `handlers` consist http handlers
- `auth` has some basic naive auth helpers
- `db` is the database layer

##### Additionally:
- I have added all api endpoint requests to the postman collection in addition to that, i have added edge cases, where we pass on wrong data, to 
  view what happens in those cases.
- I have written connection verification tests for http handlers.
- I have written tests for the database functions.

