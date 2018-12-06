package service

import (
	"log"
	"net/http"

	"github.com/BigBrother3/server/database/database"

	"github.com/unrolled/render"
)

//handle a request with method GET and path "/api/".
func filmsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("sha"))
	}
}

func getFilmsById(w http.ResponseWriter, req *http.Request) {
	/*

		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Fatal(id)
		}
	*/
	w.WriteHeader(http.StatusOK)
	log.Println(database.GetValue([]byte("film"), []byte("1")))
	w.Write([]byte("shao"))

}
