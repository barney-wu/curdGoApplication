package routes

import (
	"net/http"

	"../handlers"
	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.IndexHandler,
	},
	Route{
		"Create",
		"GET",
		"/create",
		handlers.CreateHandler,
	},
	Route{
		"Insert",
		"POST",
		"/insert",
		handlers.InsertHandler,
	},
	Route{
		"Read",
		"GET",
		"/read",
		handlers.ReadHandler,
	},
	Route{
		"Edit",
		"GET",
		"/edit",
		handlers.EditHandler,
	},
	Route{
		"Update",
		"POST",
		"/update",
		handlers.UpdateHandler,
	},
	Route{
		"Delete",
		"GET",
		"/delete",
		handlers.DeleteHandler,
	},
}

//NewRouter initialize all routers
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandleFunc)
	}
	return router
}
