package main

const (
	issuer string = "494c06f86f086f5bb135f241bada2d5ba0cd7a0d99ddcd9023b3e1eea995fa54"
	ppsKey string = "22842213df513efee733b10960b6bf19229c7a5a591f39e7cbacd18010aa537d"
	apiKey string = "22842213df513efee733b10960b6bf19229c7a5a591f39e7cbacd18010aa537d"
	apiToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDI0LTA0LTE4VDEyOjM3OjU4LjU2NDEzMjE3NFoiLCJpYXQiOjE3MTM0NDAyNzgsImlzcyI6IjQ5NGMwNmY4NmYwODZmNWJiMTM1ZjI0MWJhZGEyZDViYTBjZDdhMGQ5OWRkY2Q5MDIzYjNlMWVlYTk5NWZhNTQifQ.6J6AxTNCTLOJWpJsDyXkAkUNc0Yrv9l40Y-g49eptow"
	apiAddr string = "http://tvoykostanay.kz/api"
)

const (
	dsn string = "postgresql://postgres:52ed011d1b3d3026186e9d43ccd3011a4421ec0c65ac0f7a2893026351e64f90@localhost/work3"
	redisAddr string = "127.0.0.1:6379"
	redisPsswd string = ""
	redisDB int = 0
)

var (
	secretKey = []byte("22842213df513efee733b10960b6bf19229c7a5a591f39e7cbacd18010aa537d")
)
