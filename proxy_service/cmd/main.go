package main

import "github.com/roxyash/kmf_testtask/proxy_service/app"

//	@title			PROXY SERVICE
//	@version		1.0
//	@description	this is service, allow user authentication

//	@host		localhost:8000
//	@BasePath	/api/v1

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	proxyService := app.NewApp()

	proxyService.Run()
}
