## Homework

In this code assignment we expect you to craft a simple solution as our requirement
below

|   | Movies  |
|---|---------|
| PK | id int Auto Increment |
|    | title char(225) NOT NULL |
|    | slug char(225) unique NOT NULL |
|    | description text NOT NULL |
|    | duration INT Auto Increment |
|    | id int(5) NOT NULL |
|    | image char(255) NOT NULL |

Build a minimal service to create, update, get, and delete movies, using this boilerplate.
- The system should receive an external GET request with movie slug, ex: titanic,
  return 200 (ok) if found and 404 (not found) if not found
- The system should receive an external POST request with create movie API
  contract, with status code 201 (created) if already successful to create. Movie slug
  must be unique
- The system should receive an external PUT request with slug as movie identifier
  alongside with update movie API contract and status code 200 (ok) if successfully
  update movie content, 400 (bad request) if there is a client error, 422
  (unprocessable entity) if there’s a duplicate movie slug
- The system should receive an external DELETE request with slug as movie identifier
  return 200 (ok) if found and success to delete and 404 (not found) if the movie not
  found