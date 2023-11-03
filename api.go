package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"github.com/witczakxd/price-fetcher/types"
)

type APIFunc func (context.Context,http.ResponseWriter,*http.Request) error

type JSONAPIServer struct {
	listenAddr string
	svc PriceFetcher
}

func NewJSONApiServer(listenAddr string,svc PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc: svc,
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/",makeHTTPHandlerFunc(s.handleFetchPrice))

	http.ListenAndServe(s.listenAddr,nil)
}

func makeHTTPHandlerFunc(apiFn APIFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := context.WithValue(r.Context(), "requestID", rand.Intn(10000000))

        if err := apiFn(ctx, w, r); err != nil {
            writeJSON(w, http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
        }
    }
}


func (s *JSONAPIServer) handleFetchPrice(ctx context.Context,w http.ResponseWriter,r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price,err := s.svc.FetchPrice(ctx,ticker)


	if err != nil {
		return err
	}
	priceResp := types.PriceResponse {
		Ticker: ticker,
		Price: price,
		ReqID: ctx.Value("requestID").(int),
	}

	return writeJSON(w,http.StatusOK,&priceResp)

}

func writeJSON(w http.ResponseWriter,s int,v any) error {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}