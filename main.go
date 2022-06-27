package main

import (
	"net/http"

	"github.com/Arceister/ice-house-news/lib"
	"github.com/Arceister/ice-house-news/utils"
)

func main() {
	r := utils.NewRequestHandler()

	env := lib.NewEnv()

	appPort := env.APPPort

	http.ListenAndServe(appPort, r.Chi)
}
