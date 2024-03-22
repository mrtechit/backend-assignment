# backend-assignment
Project exposes rest endpoint to query ethereum blockchain

## API
There are 3 API's :
1. Get last parsed block
    - **URL:** `/v1/api/block`
    - **Method:** `GET`
    - **Request Body:**
    - **Response:**
        - 200 OK - Success
           ```json
          {
          "current_block": "0x4b7"
          }
           ```
        - 500 Internal Server Error
           ```json
          {
          "error": "internal error"
          }
           ```
          Other standard HTTP response codes are also possible outcomes !
    - Example curl :
      `curl --location 'http://127.0.0.1:8080/v1/api/block'`

2. Address subscription
    - **URL:** `/v1/api/subscribe/{address}`
    - **Method:** `PUT`
    - **Request Body:**
        - address (string)
    - **Response:**
        - 200 OK - Success
        - 500 Internal Server Error
          Other standard HTTP response codes are also possible outcomes !
    - Example curl :
      `curl --location 'http://127.0.0.1:8080/v1/api/subscribe/0xb794f5ea0ba39494ce839613fffba74279579268'`

3. List of transactions for an address
    - **URL:** `/v1/api/transactions`
    - **Method:** `GET`
    - **Request Body:**
    - **Response:**
        - 200 OK - Success
           ```json
          {
          "current_block": "0x4b7"
          }
           ```
        - 500 Internal Server Error
           ```json
          {
          "error": "internal error"
          }
           ```
          Other standard HTTP response codes are also possible outcomes !
    - Example curl :
      `curl --location 'http://127.0.0.1:8080/v1/api/transactions?addresss=0xb794f5ea0ba39494ce839613fffba74279579268'`

## Install Go
This project is implemented using Go. If Go is not yet installed, please download and install from [here](https://golang.org/doc/install)

## Testing:
go test ./...

## Build Server
go build -o txparser cmd/txparser/main.go

## Run server:
./txparser

## Roadmap improvement :
-Improve error handling
-Have pagination support for transactions API
-Improve logs and setups metrics
-Add validation and sanitize inputs