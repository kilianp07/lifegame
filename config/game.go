package config

type Game struct {
	Lenght      int    `json:"lenght"`
	Height      int    `json:"height"`
	Generations int    `json:"generations"`
	CharTrue    string `json:"char_true"`
	CharFalse   string `json:"char_false"`
	Sleep       string `json:"sleep"`
}
