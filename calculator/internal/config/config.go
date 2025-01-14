package config

type AppConfig struct {
	AppName string
}

func LoadConfig() AppConfig{
	return AppConfig{
		AppName: "GoCal",
	}
}