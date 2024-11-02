package main

import (
	"context"
	"github.com/OddEer0/partners-catalog/internal/infra/store/inmem"
	grpc "github.com/OddEer0/partners-catalog/internal/transport/grpc/v1"
	v1 "github.com/OddEer0/partners-catalog/protogen"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	db := inmem.NewPartnersCatalog()
	srv := grpc.NewPartnersCatalogServer(db)
	mux := runtime.NewServeMux()
	err := v1.RegisterPartnersCatalogServiceHandlerServer(ctx, mux, srv)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("starting server")
	err = http.ListenAndServe(":10100", mux)
	if err != nil {
		log.Fatal(err)
	}
}
