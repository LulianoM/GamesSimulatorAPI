package structs

type Jokenpo struct {
	UserChoice     string `json:"userChoice"`
	ComputerChoice string `json:"computerChoice"`
	YouWin         bool   `json:"youWin"`
}
