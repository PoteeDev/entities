package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/PoteeDev/admin/api/database"
	"github.com/PoteeDev/auth/middleware"
	"github.com/PoteeDev/entities/handlers"
	"github.com/gin-gonic/gin"
)

// func init() {
// 	database.ConnectDB()
// }

func main() {
	database.ConnectDB()

	appAddr := ":" + os.Getenv("PORT")
	var router = gin.Default()

	router.POST("/registration", handlers.CreateEntity)
	router.GET("/info", middleware.TokenAuthMiddleware(), handlers.GetEntityInfo)
	router.POST("/vpn", middleware.TokenAuthMiddleware())
	router.GET("/services", middleware.TokenAuthMiddleware(), handlers.GetServices)

	srv := &http.Server{
		Addr:    appAddr,
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
