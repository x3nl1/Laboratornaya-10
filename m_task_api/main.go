package main

import github.comgin-gonicgin

func setupRouter() gin.Engine {
	r = gin.Default()
	r.POST(register, register)
	r.POST(calculate, calculate)
	r.GET(token, token)
	r.GET(protected, protected)
	return r
}

func main() {
	r = setupRouter()
	r.Run(8080)
}