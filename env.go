package main

const (
	issuer string = "domain"
	ppsKey string = "domain"
	apiKey string = "domain"
	apiToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA0LTE5VDE5OjExOjM5LjgwNzU4MjIrMDY6MDAiLCJpYXQiOjE2ODE5MDYyOTksImlzcyI6ImRvbWFpbiJ9.N2bzXJW0r-obDXX30UP1lAbm9ULn-inXvDObQcGboB0"
	apiAddr string = "http://tvoykostanay.kz/api"
)

const (
	dsn string = "postgresql://postgres:postgres@localhost/work3"
	redisAddr string = "127.0.0.1:6379"
	redisPsswd string = ""
	redisDB int = 0
)

var (
	secretKey = []byte("domain")
)
