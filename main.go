package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/andrewariza/whoisServer/controllers"
	"github.com/andrewariza/whoisServer/utils"
	"github.com/fatih/color"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(time.Minute))

	r.Get("/", controllers.Record)

	r.Get("/{domain}", controllers.Whois)

	warning := color.New(color.FgRed).SprintFunc()
	if _, err := strconv.Atoi(os.Getenv("port")); err != nil {
		log.Fatal(warning("The port is invalid:"), err)
	}

	utils.Preload()
	log.Fatal(warning(http.ListenAndServe(":"+os.Getenv("port"), r)))
}
