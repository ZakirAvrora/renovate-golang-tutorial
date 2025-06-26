package main

import (
	"context"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	firestore.NewClientWithDatabase(context.Background(), "", "")
	storage.NewClient(context.Background(), storage.WithJSONReads())
	trace.NewTracerProvider()
	otel.Tracer("example.com/basic")
}
