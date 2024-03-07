# Cars-API

## Basic Information
This API provides basic functionality for creating and retrieving cars data including thier name, make, model, color, speed and type.<br>
It also provides a basic functionality for creating and retrieving cars types.<br>

## Project Structure
This project is structured based on the clean architecture.

## How to run the project
- Make sure that you have postgresSQL database called cars
- Create `.env` file that contains the following variables

```
API_HOST=?
API_PORT=?

DB_HOST=?
DB_PORT=?
DB_USER=?
DB_PASSWORD=?
DB_NAME=?
```

- Then in the main directory run the following command `go run main.go`

## How to use the API
Here is the postman documentation for the API https://documenter.getpostman.com/view/28637212/2sA2xe3ZTN

## Important Notes
- This project took total of 6 hours to be completed 2 of them in searching for the suitable packages.
- There are some imporvements that can be added to this project such as:
    - More robust error messages.
    - Pagination.