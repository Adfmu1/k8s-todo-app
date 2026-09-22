package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/nmrshll/go-cp"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	// get env vars
	godotenv.Load()
	port := ":" + os.Getenv("PORT")

	for i := 0; i < 10; i++ {
		err := cp.CopyFile(fmt.Sprintf("./static/images/image%d.jpg", i), fmt.Sprintf("/etc/kube/images/image%d.jpg", i))
		if err != nil {
			logger.Error(err.Error(), "imageNumber", i)
		}
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))

	// handle main page
	mux.Handle("/", fs)
	mux.Handle("/cache/images/", http.StripPrefix("/cache/images/", http.FileServer(http.Dir("/etc/kube/images"))))

	logger.Info("Started a server", "PORT", port)

	// create server
	serv := &http.Server{
		Addr:              port,
		Handler:           mux,
		ReadHeaderTimeout: time.Second * 5,
	}

	err := serv.ListenAndServe()
	if err != nil {
		logger.Info(err.Error())
	}
}
