package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theRemix/go90s"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/random", randomHandler)
	r.HandleFunc("/names/{key}", namesHandler)

	return r
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(go90s.GetRandomName()))
}

func namesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	w.Write([]byte(go90s.GetName(key)))
}
