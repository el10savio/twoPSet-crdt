package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// List is the HTTP handler used to return
// all the values present in the TwoPSet node in the server
func List(w http.ResponseWriter, r *http.Request) {
	// Sync the TwoPSets if multiple nodes
	// are present in a cluster
	if len(GetPeerList()) != 0 {
		TwoPSet, _ = Sync(TwoPSet)
	}

	// Get the values from the TwoPSet
	set := TwoPSet.List()

	// DEBUG log in the case of success
	// indicating the new TwoPSet
	log.WithFields(log.Fields{
		"set": set,
	}).Debug("successful twopset list")

	// JSON encode response value
	json.NewEncoder(w).Encode(set)
}
