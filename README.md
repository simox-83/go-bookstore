# CRUD API with MySQL Database

This simple example shows how to create an API, retrieving and storing data from a MySQL database.

At high level, the API comprehends 5 routes:

/book/ -> GET Method -> returns a Json with the list of all books
/book/id -> GET Method -> returns a Json with the book with id=id
/book -> POST Method -> returns a Json with the book created
/book/id -> DELETE Method -> returns a Json of the deleted book with id=id
/book/id -> PUT Method -> returns a Json with the book updated with id=id

The example uses gorilla/mux library for the API routing and the Gorm library to interact with MySQL.

Workflow: 
User curls an endpoint
The routes package registers the routes
The controller package implements the route Handlers 
The config package establishes a connection to the database and returns a db object 
The models package defines a Book object and communicates with the database to retrieve/update/delete data
The utils package is needed to perform the unMarshaling of the Json object so the database can actually perform actions on it 
The main package starts the server on port 9010

TODO: check edge cases - wrong ID, wrong payload etc..