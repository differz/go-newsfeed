package gin

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/VitaliiHurin/go-newsfeed/app/nf-service-api/api"
	"github.com/VitaliiHurin/go-newsfeed/app/nf-service-api/server"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

type Server struct {
	router *gin.Engine
}

func New(mode server.Mode, a *api.API) server.NewsfeedServer {
	router := gin.New()
	router.Use(
		Logger(),
		gin.Recovery(),
		cors.Middleware(
			cors.Config{
				Origins:        "*",
				Methods:        "GET, POST, OPTIONS, DELETE",
				RequestHeaders: "Origin, Content-Type",
			},
		),
		errorHandler(),
	)

	switch mode {
	case server.ModeRelease:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	g := router.Group("/api/v1")

	g.GET("/services", handleGetServices(a))

	g.POST("/registration", handlePostRegistration(a))
	//g.POST("/login", handlePostLogin(a))
	//g.POST("/restore-token", handlePostRestoreToken(a))
	g.GET("/user/tags", handleGetUserTags(a))
	g.POST("/user/tags", handlePostUserTags(a))
	g.DELETE("/user/tags", handleDeleteUserTags(a))

	g.GET("/articles", handleGetArticles(a))
	g.GET("/articles/:id/read", handleGetArticleMarkAsRead(a))
	g.GET("/articles/:id/unread", handleGetArticleMarkAsUnread(a))

	return &Server{
		router,
	}
}

func (s *Server) Run(httpAddr string) {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	srv := &http.Server{
		Addr:           httpAddr,
		Handler:        s.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		srv.ListenAndServe()
	}()

	<-stopChan
	fmt.Println("Shutting down API server...")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)

	fmt.Println("API server finished.")
}
