package app

import (
	"log"
	"net/http"

	"wb_go_lvl2/dev11/internal/server"
	service "wb_go_lvl2/dev11/internal/service/calendar"
	"wb_go_lvl2/dev11/internal/transport"
	"wb_go_lvl2/dev11/internal/transport/rest"
)

func Run(port string) {
	router := http.NewServeMux()
	server := new(server.Server)

	// TODO! Init Service
	calendarService := service.NewServiceCalendar()

	// TODO! Init Hadlers
	calculatorHandler := rest.NewHandlerCalendar(calendarService)

	// TODO! Call Handlers
	calculatorHandler.Init(router)

	// TODO! Check Middleware
	allowHosts := []string{}
	loggedRouter := transport.CorsMiddleware(allowHosts, router)

	// TODO! Server Run
	if err := server.Run(port, loggedRouter); err != nil {
		log.Fatalf("ошибка при запуске сервера: %v", err)
		return
	}
}
