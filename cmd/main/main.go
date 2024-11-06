package main

import (
	"context"
	"github.com/OddEer0/partners-catalog/internal/infra/store/inmem"
	grpcv1 "github.com/OddEer0/partners-catalog/internal/transport/grpc/v1"
	v1 "github.com/OddEer0/partners-catalog/protogen"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(l)
	ctx := context.Background()
	db := inmem.NewPartnersCatalog()
	srv := grpcv1.NewPartnersCatalogServer(db)
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
