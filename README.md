# Go-server (My first Go project)

Welcome to My First Go Project, a simple API with CRUD (Create, Read, Update, Delete) functionality built using Go programming language. This API allows you to manage a list of movies, including creating new movies, reading movie details, updating existing movies, and deleting movies.

## How to run the API

1. Clone the repository to your local machine using the following command:

    ```$ git clone https://github.com/[username]/my-first-go-project.git```

2. Change the working directory to the cloned repository:

    ```cd Go-server```

3. Build and run the API using the following command:

    ```$ git clone https://github.com/[username]/my-first-go-project.git```

4. The API should now be running on **http://localhost:8000/**. You can access the API using a tool like curl or a web browser.

## API Endpoints
The API has the following endpoints:

- GET /movies: Returns a list of all movies stored in the API.
- GET /movies/{id}: Returns the details of a movie with the specified ID.
- POST /movies: Creates a new movie using the information provided in the request body.
- PUT /movies/{id}: Updates an existing movie with the specified ID using the information provided in the request body.
- DELETE /movies/{id}: Deletes a movie with the specified ID.
