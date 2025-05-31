package config

type Config struct {
	Server Server
	App    App
}

type App struct {
	Name string
}

type Server struct {
	Host string
	Port string
}
