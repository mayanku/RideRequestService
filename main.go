package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"riderequest/handler"
	"riderequest/initializers"
	"time"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()
	l := log.New(os.Stdout, "products-api ", log.LstdFlags) //for logging

	router.GET("/slottable", handler.SlotTable)
	router.POST("/bookingRequest", handler.BookingRequest)
	router.POST("/saveToTaskQueue", handler.SaveToTaskQueue)
	//router.GET("/CabAllocation",handler.CabAllocation)

	address := ":8080"          // Set the desired address
	timeout := 15 * time.Second // Set the desired timeout duration
	handler := http.TimeoutHandler(router, timeout, "Request timed out")

	server := &http.Server{
		Addr:         address,
		Handler:      handler,
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// Handle the error if the server failed to start
		// For example, log the error or exit the application
		panic(err)
	}

	// start the server
	go func() {
		l.Println("Starting server on port 8080")

		err := server.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	//Product handler
	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)

}
