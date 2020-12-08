package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// IsPresent is the JSON struct
// encapsulating the Lookup Response
type IsPresent struct {
	Present bool `json:"present"`
}

// Lookup is the HTTP handler used to return
// if a given value is present in the TwoPSet node in the server
func Lookup(w http.ResponseWriter, r *http.Request) {
	var err error
	var present bool

	// Obtain the value from URL params
	value := mux.Vars(r)["value"]

	// Sync the TwoPSets if multiple nodes
	// are present in a cluster
	if len(GetPeerList()) != 0 {
		TwoPSet, _ = Sync(TwoPSet)
	}

	// Lookup given value in the TwoPSet
	present, err = TwoPSet.Lookup(value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to lookup twopset value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DEBUG log in the case of success indicating
	// the new TwoPSet, the lookup value and if its present
	log.WithFields(log.Fields{
		"set":     TwoPSet,
		"value":   value,
		"present": present,
	}).Debug("successful twopset lookup")

	isPresent := IsPresent{present}

	JSONResponse, err := json.Marshal(isPresent)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to json marshall lookup twopset value")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(JSONResponse)
	w.WriteHeader(http.StatusOK)
}
