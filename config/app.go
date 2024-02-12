package config

type App struct {
	DBConf DBConf `json:"db"`
	Game   Game   `json:"game"`
}
