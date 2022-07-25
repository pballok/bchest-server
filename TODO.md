*** GUIDES ***
https://www.howtographql.com/graphql-go/1-getting-started/

Example with Docker compose and PostgreSQL:
https://servian.dev/building-a-graphql-api-in-go-using-gqlgen-f7a42eba2193

Example for GraphQL unit test:
https://github.com/99designs/gqlgen/blob/master/_examples/todo/todo_test.go


*** DEVELOPMENT ***
start with: go run server.go

send requests to localhost:8080/query
Example request:
query {
    listCharacters (player: "Player1") {
        name
        description
    }
}
