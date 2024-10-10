# Graph sql Service - Integrate in Echo Framework
This project demonstrates a simple Go application with GraphQL which is integrated Echo Framework, running inside a Docker container and managed with Docker Compose.

## Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Structure design

```
|-- cmd
    |-- server
        |-- main.go (entry point for the application)
|-- internal
    |-- app
        |-- app.go (Echo instance and setup)
        |-- middleware (custom middleware for the Echo instaenc)
        |-- routes (definition of all application routes)
    |-- domain
        |-- model (structs for domain objects)
        |-- repository (interfaces for data access)
        |-- service (interfaces for business logic)
        |-- usecase
    |-- infrastructure
        |-- persistence (implementations of repositories)
        |-- utils (utility functions)
|-- pkg (third-party packages)
    |-- logger (implementations of logging lib)
        |-- logger.go
    |-- error (definition errors for the project)
        |-- error.go
```

In this structure **`internal`** directory contains all the code specific to the
application, including the **`app`**, **`domain`**, **`infrastructure`**
packages. The **`cmd`** directory contains the **`main.go`** file, which is the
entry point of the application. The **`pkg`** directory contains third-party
packages used in this application.

## Setup and Run

### 1. Clone the Repository

```sh
git clone https://github.com/JJitusDD/GraphSql.git
cd GraphSql
docker-compose up --build
```

``` 
# Make a POST request to create a new Todo
curl --location 'localhost:8081/query' \
 --header 'Content-Type: application/json' \
 --header 'Authorization: ••••••' \

	--data '{
	   "query": "mutation { createTodo(input: {userId: \"id1\", text: \"john.doe\"}) { id text done user { id name } } }"
	}'

# Make a POST request to get list Todo	
curl --location 'http://localhost:8081/query' \
--header 'Content-Type: application/json' \
--data '{"query":"{ todos { id text done user { id name } } }"}'	
```