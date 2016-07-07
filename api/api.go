package api

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"github.com/NebulousLabs/Sia/encoding"
	"github.com/robvanmieghem/siapool/sharechain"
)

//PoolAPI implements the http handlers
type PoolAPI struct {
	//Fee is the poolfee in 0.01%
	Fee int
	//ShareChain for getting work and posting shares
	ShareChain *sharechain.ShareChain
}

//FeeHandler writes the fee applied by the pool
func (pa *PoolAPI) FeeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%.2f%%", float64(pa.Fee)/100)
}

//GetWorkHandler returns a header to the miners to work on
func (pa *PoolAPI) GetWorkHandler(w http.ResponseWriter, r *http.Request) {
	//Make sure the response does not get cached somewhere
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	w.Header().Set("Expires", "0")                                         // Proxies.

	payoutaddress := mux.Vars(r)["payoutaddress"]
	log.Debugln("GetWork from", payoutaddress)

	//TODO: check if the request was not made too fast after the previous one

	bhfw, target, err := pa.ShareChain.HeaderForWork(payoutaddress)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(encoding.MarshalAll(target, bhfw))
}

//SubmitHeaderHandler is called by the miners to submit their shares
// A 204 is returned when successful
func (pa *PoolAPI) SubmitHeaderHandler(w http.ResponseWriter, r *http.Request) {
	payoutaddress := mux.Vars(r)["payoutaddress"]
	log.Debugln("Processing headersubmission from", payoutaddress)

	//TODO: reset the request for new header timeout
}
