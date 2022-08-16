# BattleChest server

*** RUNNING ***
start with: go run server.go

send requests to localhost:8080/query
Example request:
query {
    listCharacters (player: "Player1") {
        name
        description
    }
}

*** TESTING ***
go test ./...
