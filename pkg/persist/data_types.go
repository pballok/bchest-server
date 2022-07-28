package persist

type PlayerData struct {
	Name           string `json:"name"`
	HashedPassword string `json:"password"`
}

type CharacterData struct {
	Name        string `json:"name"`
	PlayerName  string `json:"player_name"`
	Description string `json:"description"`
}
