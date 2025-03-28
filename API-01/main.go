package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"server/server"
	"syscall"
)

func main() {
	ctx := context.Background()

	serverDoneChan := make(chan os.Signal, 1)
	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	srv := server.New(":8080")

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println("Error al iniciar el servidor:", err)
		}
	}()

	log.Println("Servidor iniciado en :8080")
	

	<- serverDoneChan

	srv.Shutdown(ctx)
	log.Println("Servidor detenido")
}
