# json-to-html-server
Simple web application in the Go programming language that performs JSON-to-HTML conversion.

## Constraints
Server accepts date only in the following layout `"2019-04-10"`.

## Run server
Execute `make run`. Server will then run with default arguments.

Arguments:
- You can specify `port` with switch `-p`
- Path to main page HTML with `-m`
- Path to threads template with `-t` 