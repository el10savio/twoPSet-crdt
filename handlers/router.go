package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/el10savio/twoPSet-crdt/twopset"
)

var (
	// TwoPSet is the 2PSet
	// data structure initialized
	TwoPSet twopset.TwoPSet
)

func init() {
	TwoPSet = twopset.Initialize()
}

// Route defines the Mux
// router individual route
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Routes is a collection
// of individual Routes
var Routes = []Route{
	{"/", "GET", Index},
	{"/twopset/list", "GET", List},
	{"/twopset/values", "GET", Values},
	{"/twopset/lookup/{value}", "GET", Lookup},
	{"/twopset/add/{value}", "POST", Add},
	{"/twopset/remove/{value}", "POST", Remove},
}

// Index is the handler for the path "/"
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World TwoPSet Node\n")
}

// Logger is the middleware to
// log the incoming request
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"path":   r.URL,
			"method": r.Method,
		}).Info("incoming request")

		next.ServeHTTP(w, r)
	})
}

// Router returns a mux router
func Router() *mux.Router {
	router := mux.NewRouter()

	for _, route := range Routes {
		router.HandleFunc(
			route.Path,
			route.Handler,
		).Methods(route.Method)
	}

	router.Use(Logger)

	return router
}
