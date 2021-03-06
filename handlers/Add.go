package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Add is the HTTP handler used to append
// values to the TwoPSet node in the server
func Add(w http.ResponseWriter, r *http.Request) {
	var err error

	// Obtain the value from URL params
	value := mux.Vars(r)["value"]

	// Add the given value to our stored TwoPSet
	TwoPSet, err = TwoPSet.Addition(value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to add value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DEBUG log in the case of success indicating
	// the new TwoPSet and the value added
	log.WithFields(log.Fields{
		"set":   TwoPSet,
		"value": value,
	}).Debug("successful twopset addition")

	// Return HTTP 200 OK in the case of success
	w.WriteHeader(http.StatusOK)
}
