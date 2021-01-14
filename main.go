package main

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var chiLambda *chiadapter.ChiLambda

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}

func show(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("show"))
}

func store(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("store"))
}

func update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update"))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if chiLambda == nil {
		time.Sleep(10 * time.Millisecond)
		r := chi.NewRouter()
		r.Use(render.SetContentType(render.ContentTypeJSON))
		r.Route("/recipes", func(r chi.Router) {
			r.Get("/", index)
			r.Post("/", store)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", show)
				r.Put("/", update)
				r.Delete("/", delete)
			})
		})
		chiLambda = chiadapter.New(r)
	}
	return chiLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
