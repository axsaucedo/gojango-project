module github.com/axsaucedo/gjango-myapp

go 1.18

replace github.com/axsaucedo/gjango => ../gjango

require (
	github.com/axsaucedo/gjango v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.7
)

require github.com/joho/godotenv v1.4.0 // indirect
