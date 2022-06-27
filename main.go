package main

import (
	"net/http"

	"github.com/Arceister/ice-house-news/utils"
)

func main() {
	r := utils.NewRequestHandler()
	r.Chi.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "Api Hit!"}`))
	})
	http.ListenAndServe(":5055", r.Chi)
}
