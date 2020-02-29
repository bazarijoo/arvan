package http

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("GET").Path("/balance/{phone_number}").Handler(httptransport.NewServer(
		endpoints.GetBalance,
		decodeGetBalanceReq,
		encodeResponse,
	))

	r.Methods("POST").Path("/update-balance").Handler(httptransport.NewServer(
		endpoints.UpdateBalance,
		decodeUpdateBalanceReq,
		encodeResponse,
	))

	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
