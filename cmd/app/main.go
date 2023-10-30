package main

import (
	"boardType/cmd/app/handler"
	"boardType/internal/controller"
	"boardType/internal/repository"
	"boardType/internal/repository/infla"
	"boardType/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	h := getHandler()

	r.PathPrefix("/board-types").Handler(h)
	http.ListenAndServe(":8085", r)
}

func getHandler() http.Handler {
	return handler.NewHandler(controller.NewBoardTypeController(service.NewBoardTypeService(repository.NewBoardTypeRepository(infla.NewDB()))))
}
