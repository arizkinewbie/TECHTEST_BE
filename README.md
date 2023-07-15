# Technical Test – Backend Engineer

## About

Cake Store RESTFul API Service within the specified time

## Build With

* [Go](https://go.dev/)
* [MySQL - db4free.net](https://db4free.net/)
* [Google App Engine (GAE)](https://cloud.google.com/appengine)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/arizkinewbie/TECHTEST_BE.git
   ```

2. Navigate to the project directory:

   ```bash
   cd TECHTEST_BE
   ```

3. Install the dependencies:

   ```bash
   go mod tidy
   ```

4. Set up the database:
   * Create a MySQL database for the application.
   * Update the database configuration in the `.env` file.

5. Run the application:

   ```bash
   go run main.go
   ```

6. The API will be accessible at `http://localhost:8000`.

## Usage

Use a tool like Postman to send HTTP requests to the API endpoints.
| Method | Endpoint          | Description                     |
| ------ | ----------------- | ------------------------------- |
| GET    | /cakes            | Get all cakes                   |
| GET    | /cakes/{id}       | Get a cake by ID                |
| POST   | /cakes            | Add a new cake                  |
| PUT    | /cakes/{id}       | Update a cake                   |
| DELETE | /cakes/{id}       | Delete a cake                   |

## Documentation

| API Name | Link |
| ------ | ------ |
| Cake Store RESTFul API | [![Run in Postman](https://run.pstmn.io/button.svg)](https://documenter.getpostman.com/view/27407764/2s946fdYJa) |
Deployment URL : <https://steam-mantis-381321.et.r.appspot.com/>
