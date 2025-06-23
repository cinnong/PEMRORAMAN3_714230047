package config

var allowedOrigins = []string{
	"http://localhost:3000",
	"https://indrariksa.github.io",
	"http://localhost:5173/",
	"http://127.0.0.1:8088/",
}

func GetAllOrigins() []string {
	return allowedOrigins
}