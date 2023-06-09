package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
)

// RegisterRoutes registers minting module REST handlers on the provided router.
func RegisterRoutes(clientCtx client.Context, rtr *mux.Router) {
	r := WithHTTPDeprecationHeaders(rtr)
	registerQueryRoutes(clientCtx, r)
}

// DeprecationURL is the URL for migrating deprecated REST endpoints to newer ones.
// TODO Switch to `/` (not `/master`) once v0.40 docs are deployed.
// https://github.com/cosmos/cosmos-sdk/issues/8019
const DeprecationURL = "https://docs.cosmos.network/master/migrations/rest.html"

// addHTTPDeprecationHeaders is a mux middleware function for adding HTTP
// Deprecation headers to a http handler
func addHTTPDeprecationHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Deprecation", "true")
		w.Header().Set("Link", "<"+DeprecationURL+">; rel=\"deprecation\"")
		w.Header().Set("Warning", "199 - \"this endpoint is deprecated and may not work as before, see deprecation link for more info\"")
		h.ServeHTTP(w, r)
	})
}

func WithHTTPDeprecationHeaders(r *mux.Router) *mux.Router {
	subRouter := r.NewRoute().Subrouter()
	subRouter.Use(addHTTPDeprecationHeaders)
	return subRouter
}
