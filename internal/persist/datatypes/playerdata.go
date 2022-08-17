package datatypes

type PlayerData struct {
	Name           string `json:"name"`
	HashedPassword string `json:"password"`
}
