package routes

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"test/database"
	"test/handler"
	"test/middleware"
	"test/service"
)

var validate = validator.New()
var sellerHandler = handler.NewSellerHandler(service.NewSellerService(database.NewSellerDatabase(), validate))

func UserRoutes(mux *http.ServeMux, username, password string) {
	mux.Handle("/sellers", middleware.BasicAuth(username, password, http.HandlerFunc(sellerHandler.GetAll)))
	mux.Handle("/sellers/id", middleware.BasicAuth(username, password, http.HandlerFunc(sellerHandler.GetByID)))
	mux.Handle("/sellers/create", middleware.BasicAuth(username, password, http.HandlerFunc(sellerHandler.Post)))
	mux.Handle("/sellers/update", middleware.BasicAuth(username, password, http.HandlerFunc(sellerHandler.Put)))
	mux.Handle("/sellers/delete", middleware.BasicAuth(username, password, http.HandlerFunc(sellerHandler.Delete)))
}
