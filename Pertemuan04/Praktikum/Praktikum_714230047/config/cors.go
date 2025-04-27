package config

var allowedOrigins = []string{
	"http://localhost:3000",
	"https://indrariksa.github.io",
}

func GetAllOrigins() []string {
	return allowedOrigins
}