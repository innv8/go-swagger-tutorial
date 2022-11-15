package controllers

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func (b *Base) DocsHandler() http.HandlerFunc {
	return httpSwagger.Handler(
		httpSwagger.URL("http://localhost:5000/docs/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)
}
