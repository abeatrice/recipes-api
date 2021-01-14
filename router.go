package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/events"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var chiLambda *chiadapter.ChiLambda

// Router ...
func Router(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if chiLambda == nil {
		time.Sleep(10 * time.Millisecond)
		r := chi.NewRouter()
		r.Use(render.SetContentType(render.ContentTypeJSON))
		r.Route("/recipes", func(r chi.Router) {
			r.Get("/", Index)
			r.Post("/", Store)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", Show)
				r.Put("/", Update)
				r.Delete("/", Delete)
			})
		})
		chiLambda = chiadapter.New(r)
	}
	return chiLambda.ProxyWithContext(ctx, req)
}
