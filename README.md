# FizzBuzz API

This project implements a REST API endpoint for the FizzBuzz problem. Given two strings (string1 and string2) and three integers (int1, int2, and limit), the API returns a list of strings with numbers from 1 to limit where:

- All multiples of int1 are replaced by string1
- All multiples of int2 are replaced by string2
- All multiples of both int1 and int2 are replaced by a string1 and string2 concatenation


## Getting started

### Run on your local machine

To run the FizzBuzz API on your locale machine, run the command `make run`. Application will be run on port `8090`.

The FizzBuzz endpoint should now be accessible at `http://localhost:8090/fizzbuzz`.

### Run in Docker

To run the FizzBuzz API with Docker, you'll need to have Go and Docker installed. Follow these steps:

1. Clone the repository to your local machine.
2. Build the Docker image by running `make docker-build DOCKER_TAG={tag_name}`
3. Run the Docker container by running `make docker-run DOCKER_PORT={port_number} DOCKER_TAG={tagname}`

By default `DOCKER_PORT` will be equal to **8090** and `DOCKER_TAG` will be equal to **latest**.

The FizzBuzz endpoint should now be accessible at `http://localhost:{port_number}/fizzbuzz`.


## Usage

To use the FizzBuzz API, send a GET request to the /fizzbuzz endpoint with a JSON payload containing the following fields:

- int1: an integer that represents the first multiple.
- int2: an integer that represents the second multiple.
- string1: a string that replaces int1 in the output when the number is a multiple of int1.
- string2: a string that replaces int2 in the output when the number is a multiple of int2.
- limit: an integer that represents the upper limit of the output list. 

### Example

    curl --location --request GET 'http://localhost:8090/fizzbuzz' \
        --header 'Content-Type: application/json' \
        --data-raw '{
        "string1": "fizz",
        "string2": "buzz",
        "int1": 3,
        "int2": 5,
        "limit": 15
    }'

The response will be a string such as :

    "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz"

### Tests

To run unit tests use command `make test`.

### Lint

To run linter use command `make lint`.

Note that  Docker is required as linter is run within a Docker container.