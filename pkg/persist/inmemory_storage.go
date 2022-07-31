package persist

type inMemoryPlayerStorage struct {
	inMemoryTable[string, PlayerData]
}

type inMemoryCharacterStorage struct {
	inMemoryTable[string, CharacterData]
}

type inMemoryStorageType struct {
	players    inMemoryPlayerStorage
	characters inMemoryCharacterStorage
}

func (s *inMemoryStorageType) Init() bool {
	// Temporary init code below
	tempPlayer := PlayerData{
		Name:           "pballok",
		HashedPassword: "hash",
	}
	s.players.AddNew(tempPlayer.Name, &tempPlayer)

	return true
}

func (s *inMemoryStorageType) Players() playerStorage {
	return s.players
}

func (s *inMemoryStorageType) Characters() characterStorage {
	return s.characters
}

var inMemoryStorage = inMemoryStorageType{
	players: inMemoryPlayerStorage{
		inMemoryTable: inMemoryTable[string, PlayerData]{
			name: "Players",
		},
	},
	characters: inMemoryCharacterStorage{
		inMemoryTable: inMemoryTable[string, CharacterData]{
			name: "Characters",
		},
	},
}
