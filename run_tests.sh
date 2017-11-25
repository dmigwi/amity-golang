# !bin/bash 

# current parent package environment variable
AG_PKG=amity-golang

# Set the ENV environment variable to mockDB
# controllers tests should run with the mock models connection
export ENV="mockDB"

# ENV_FILE exposes the location of the .env file that is located 
# at the root of folder for the test environment
export ENV_FILE="../.env"

# Run the controllers  and the models tests
go test  ${AG_PKG}/controllers  ${AG_PKG}/models $@

# Set the ENV environment variable to DB
# controllers tests should run with the db models connection
export ENV="DB"

# Run the controllers using the set environment
go test ${AG_PKG}/controllers $@