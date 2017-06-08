package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theRemix/go90s"
)

func RandomHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(go90s.GetRandomName()))
}

func NamesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	w.Write([]byte(go90s.GetName(key)))
}
