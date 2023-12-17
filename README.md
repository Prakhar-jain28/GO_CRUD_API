
# Blogging Website API

This is a Basic CRUD(Create, Read, Update, Delete) API project for a blogging Website made using GoFr framework of GoLang.


## Features

- **Post A Blog:** Inserts a new Blog in the database.
- **Get a Blog By ID:** Retrives a Blog by its ID from the database.
- **Delete a Blog:** Deletes a Blog by its ID.
- **Update A Blog:** Updates a Blog by its ID.




## Installation

To run this project on your local machine:

1. Clone this Project by using this command in the git bash terminal.

```Bash  
Git clone https://github.com/Prakhar-jain28/GO_CRUD_API.git
```

2. Open the project folder inside VSCode.
3. In the terminal type `go mod tidy` to install all the dependencies.
4. Now type `go run main.go` to run the project.



## Usage

Key Endpoints for this project: 

- **POST Blog Endpoint:** `POST /blog` 
    - Schema Used: 
```go
    //Pass the body according to this schema
    type Blog struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Content string `json:"content"`
    }
```

- **GET Blog Endpoint:** `GET /blog/{ID}` 
- **DELETE Blog Endpoint:** `DELETE /blog/{ID}` 
- **Update Blog Endpoint:** `PUT /blog/{ID}` 