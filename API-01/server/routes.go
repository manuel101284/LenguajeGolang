package server

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", index)

	http.HandleFunc("/countries", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodGet:
				getCountries(rw, r)
			case http.MethodPost:
				addCountry(rw, r)
			default:
				rw.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(rw, "Método no permitido")
				return
		}
	})

	http.HandleFunc("/countries/add", addCountry)
	
	fmt.Println("Rutas inicializadas")
}