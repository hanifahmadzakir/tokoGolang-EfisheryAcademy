package app

import "tokoGolang/app/controllers"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controller.Home).Methods("GET")
}