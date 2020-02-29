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

	r.Methods("GET").Path("/voucher-status/{voucher_code}").Handler(httptransport.NewServer(
		endpoints.GetVoucherCodeStatus,
		decodeGetVoucherCodeStatusReq,
		encodeResponse,
	))

	r.Methods("POST").Path("/submit-voucher-code").Handler(httptransport.NewServer(
		endpoints.SubmitVoucherCode,
		decodeSubmitVoucherCodeRequest,
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
