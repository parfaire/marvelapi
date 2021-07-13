package models

// Character
type CharacterDataWrapper struct {
	Data CharacterDataContainer
}

type CharacterDataContainer struct {
	Results []Character
}

type Character struct {
	Id          int
	Name        string
	Description string
}

// Character ID Only
type CharacterDataWrapperIdOnly struct {
	Data CharacterDataContainerIdOnly
}

type CharacterDataContainerIdOnly struct {
	Results []CharacterIdOnly
	Total   int
}

type CharacterIdOnly struct {
	Id int
}
