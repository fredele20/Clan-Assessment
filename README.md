# Clan-Africa-Assessment
This assesment is writted in Go (Golang) programming Language.
Mock databases in JSON format are used as data storage for this assessment, you can find them in the db folder of the project

To setup and run locally:
1. Install the Go programming language on your machine
2. Open your terminal and go to the project directory and run "go mod tidy" this will install all the dependecies locally.
3. You can run the program directly by running "go run main.go"

To test endpoints:
Download and open any REST Client of your choice (eg: Postman, Thunder Client(VS Code extention), RestFox) etc.
1. http://localhost:6000/flights/  --> This returns all the available endpoints
2. http://localhost:6000/flights/?departure=city&destination=city&date=date(optional) --> This returns all the availble flights based on the user query input
3. http://localhost:6000/flights/options/ --> This returns all the availabe flights options
