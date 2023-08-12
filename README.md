# Clan-Africa-Assessment
This assesment is writted in Go (Golang) programming Language.
Mock databases in JSON format are used as data storage for this assessment, you can find them in the db folder of the project

Mock Databases:
1. flights.json: Contains all the available flights, you can add more to the JSON file by going to PopulateFlight function in the flight.go file in the db folder
2. flightOptions.json: Contains the number of options available for flights with their respective prices. You can also edit it the PopulateFlightOption function in the flight.go file
3. bookings.json: This is where any successful booked flight information goes to. Informations are only stored here when you hit the No. 4 endpoint in the list of endpoint below

To setup and run locally:
1. Install the Go programming language on your machine
2. Open your terminal and go to the project directory and run "go mod tidy" this will install all the dependecies locally.
3. You can run the program directly by running "go run main.go"

To test endpoints:
Download and open any REST Client of your choice (eg: Postman, Thunder Client(VS Code extention), RestFox) etc.
1. http://localhost:6000/flights/  --> This returns all the available endpoints
2. http://localhost:6000/flights/?departure=city&destination=city&date=date(optional) --> This returns all the availble flights based on the user query input
3. http://localhost:6000/flights/options/ --> This returns all the availabe flights options
4. http://localhost:6000/flights/:id/book/ --> This endpoint is for booking a flight, user is required to enter a flight       option in order to successfully book a flight. The flight id to be booked is passed in place of the :id in the endpoint example: (http://localhost:6000/flights/1/book) to book the flight with the id of 1.

Unit Testing:
The test files are available in the db folder/package of this project, they all have _test.go attached to their file names.
To run all the test files, you have to change the directory to point to the db package, and type or run "go test -v" on your terminal or CMD
Example: /home/fredel/Documents/go/clan-africa/db --> This is how my working directory looks like on my linux machine
You can also test the individual test cases separately by clicking "run test" text above any test function or click on the play icon on the left.
