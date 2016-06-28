package api

import (
	"fmt"
	"net/http"

	"github.com/robvanmieghem/siapool/siadclient"
)

//PoolAPI implements the http handlers
type PoolAPI struct {
	//Fee is the poolfee in 0.01%
	Fee int
	//SiadClient is the client towards the sia daemon
	SiadClient *siadclient.SiadClient
}

//FeeHandler writes the fee applied by the pool
func (pa *PoolAPI) FeeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%.2f%%", float64(pa.Fee)/100)
}
