package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/el10savio/twoPSet-crdt/twopset"
)

// Sync merges multiple TwoPSet present in a network to get them in sync
// It does so by obtaining the TwoPSet from each node in the cluster
// and performs a merge operation with the local TwoPSet
func Sync(TwoPSet twopset.TwoPSet) (twopset.TwoPSet, error) {
	// Obtain addresses of peer nodes in the cluster
	peers := GetPeerList()

	// Return the local TwoPSet back if no peers
	// are present along with an error
	if len(peers) == 0 {
		return TwoPSet, errors.New("nil peers present")
	}

	// Iterate over the peer list and send a /twopset/values GET request
	// to each peer to obtain its TwoPSet
	for _, peer := range peers {
		peerTwoPSet, err := SendListRequest(peer)
		if err != nil {
			log.WithFields(log.Fields{"error": err, "peer": peer}).Error("failed sending twopset values request")
			continue
		}

		// Merge the peer's TwoPSet with our local TwoPSet
		TwoPSet = twopset.Merge(TwoPSet, peerTwoPSet)
	}

	// DEBUG log in the case of success
	// indicating the new TwoPSet
	log.WithFields(log.Fields{
		"set": TwoPSet,
	}).Debug("successful twopset sync")

	// Return the synced new TwoPSet
	return TwoPSet, nil
}

// SendListRequest is used to send a GET /twopset/values
// to peer nodes in the cluster
func SendListRequest(peer string) (twopset.TwoPSet, error) {
	var _twopset twopset.TwoPSet

	// Return an empty TwoPSet followed by an error if the peer is nil
	if peer == "" {
		return _twopset, errors.New("empty peer provided")
	}

	// Resolve the Peer ID and network to generate the request URL
	url := fmt.Sprintf("http://%s.%s/twopset/values", peer, GetNetwork())
	response, err := SendRequest(url)
	if err != nil {
		return _twopset, err
	}

	// Return an empty TwoPSet followed by an error
	// if the peer's response is not HTTP 200 OK
	if response.StatusCode != http.StatusOK {
		return _twopset, errors.New("received invalid http response status:" + fmt.Sprint(response.StatusCode))
	}

	// Decode the peer's TwoPSet to be usable by our local TwoPSet
	var twoPSet twopset.TwoPSet
	err = json.NewDecoder(response.Body).Decode(&twoPSet)
	if err != nil {
		return _twopset, err
	}

	// Return the decoded peer's TwoPSet
	_twopset = twoPSet
	return _twopset, nil
}
