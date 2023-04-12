# BookStore

## Prerequisites
### Install docker
This project is dockerized, so to run it, you'll need Docker and Docker Compose installed.
## Endpoints and examples

The following endpoints are available:
- `GET /api/books`: Returns a list of books, ordered by price in descending order by default. Optional parameters:
- `title`: search for books with a title that matches the given string.
- `price`: order books by price in ascending order.

Example:

```bash

$ curl http://localhost:8080/api/books?title=Python\&price=asc
{
  "status": "success",
  "data": [
    {
      "id": 1,
      "title": "Python Cookbook",
      "description": "A collection of Python recipes",
      "price": 39.99
    },
    {
      "id": 2,
      "title": "Learning Python",
      "description": "An introduction to Python programming",
      "price": 49.99
    }
  ]
}
``` 
- `GET /api/books/:id`: Returns the book with the given ID. If no book is found with the given ID, a `404` response is returned.

Example:

```bash

$ curl http://localhost:8080/api/books/1
{
  "status": "success",
  "data": {
    "id": 1,
    "title": "Python Cookbook",
    "description": "A collection of Python recipes",
    "price": 39
  }
}
``` 
- `POST /api/books`: Creates a new book with the given parameters in the request body. Required parameters:
- `title`
- `description`
- `price`

Example:

```bash

$ curl -X POST -H "Content-Type: application/json" -d '{"title": "Django for Beginners", "description": "An introduction to Django web framework", "price": 29.99}' http://localhost:8080/api/books
{
  "status": "success",
  "data": {
    "id": 3,
    "title": "Django for Beginners",
    "description": "An introduction to Django web framework",
    "price": 29
  }
}
``` 
- `PUT /api/books/:id`: Updates the book with the given ID. The request body should contain only the fields that need to be updated. If no fields are given or the book with the given ID is not found, a `400` response is returned.

Example:

```bash

$ curl -X PUT -H "Content-Type: application/json" -d '{"title": "Python Recipes", "price": 49.99}' http://localhost:8080/api/books/1
{
  "status": "success",
  "data": {
    "id": 1,
    "title": "Python Recipes",
    "description": "A collection of Python recipes",
    "price": 49
  }
}
``` 
- `DELETE /api/books/:id`: Deletes the book with the given ID. If no book is found with the given ID, a `404` response is returned.

Example:

```bash

$ curl -X DELETE http://localhost:8080/api/books/1
{
  "status": "success",
  "message": "Book deleted successfully"
}
```
## How to run the project

To run the project, follow these steps:
1. Clone the repository

2. Navigate to the project directory
3. Set up environment variables by creating a `.env` file in the root directory of the project and adding the necessary variables. You can refer to the "env vars" section of the README for more details on the required variables.
4. Build the Docker image using the following command:
```bash
sudo docker build -t bookstore-api .   
``` 
5. Run the Docker container using the following command:

```bash
sudo docker run -it --net=host bookstore-api 
```
This will start the server and make it accessible on `http://localhost:8080`.
6. You can test the API endpoints using tools like `curl` or an API client like `Postman`. Refer to the "Endpoints and Examples" section of the README for more information on the available endpoints and their usage.
