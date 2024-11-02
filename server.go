package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dddong3/Bid_Backend/auctionitem"
	"github.com/dddong3/Bid_Backend/config"
	"github.com/dddong3/Bid_Backend/graph"
	"github.com/dddong3/Bid_Backend/graph/resolvers"
	"github.com/dddong3/Bid_Backend/logger"
	"github.com/dddong3/Bid_Backend/middlewares"
	"github.com/dddong3/Bid_Backend/rest"
	"github.com/dddong3/Bid_Backend/rest/handlers"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func main() {
	port := config.GetEnv("PORT", "8080")
	defer logger.Sync()

	r := chi.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler)
	r.Use(middlewares.TimingMiddleware)

	rest.RegisterRoutes(r, &handlers.AuctionItemHandler{
		Service: &auctionitem.AuctionItemService{
			Repo: auctionitem.GetAuctionItemRepo(),
		},
	})

	resolvers := resolvers.InitResolver()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolvers}))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	logger.Logger.Infof("connect to http://0.0.0.0:%s/ for GraphQL playground", port)
	logger.Logger.Fatal(http.ListenAndServe("0.0.0.0:"+port, r))
}
