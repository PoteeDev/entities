package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/PoteeDev/auth/middleware"
	"github.com/PoteeDev/team/handlers"
	"github.com/explabs/ad-ctf-paas-api/database"
	"github.com/gin-gonic/gin"
)

func init() {
	database.InitMongo()
}

func main() {

	appAddr := ":" + os.Getenv("PORT")
	var router = gin.Default()

	router.POST("/registration", handlers.CreateTeam)
	router.GET("/info", middleware.TokenAuthMiddleware(), handlers.GetTeamInfo)
	router.POST("/vpn", middleware.TokenAuthMiddleware())

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
