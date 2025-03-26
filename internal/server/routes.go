package server

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	// Fix the import paths to use the correct module name
	graph "jobfai-analytics/internal/graph"
	resolver "jobfai-analytics/internal/graph/resolver"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Basic API routes
	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)

	// GraphQL setup
	graphqlHandler := s.setupGraphQLHandler()
	playgroundHandler := playground.Handler("GraphQL Playground", "/query")

	// GraphQL routes
	r.POST("/query", gin.WrapH(graphqlHandler))
	r.GET("/query", gin.WrapH(graphqlHandler)) // For subscriptions via WebSocket
	r.GET("/playground", gin.WrapH(playgroundHandler))

	return r
}

func (s *Server) setupGraphQLHandler() http.Handler {
	// Create a new resolver with database access
	resolvers := resolver.NewResolver(s.db)

	// Create the GraphQL server
	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolvers,
	}))

	// Configure WebSocket transport for subscriptions
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Allow all origins
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	// Configure other transports
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	// Add extensions
	srv.Use(extension.Introspection{})
	// srv.Use(extension.AutomaticPersistedQuery{
	// 	Cache: lru.New(100),
	// })

	return srv
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
