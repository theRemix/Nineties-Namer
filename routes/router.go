package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/theRemix/go90s"
)

const STATIC_DIR = "./public/"

func Router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/random", randomHandler)
	r.HandleFunc("/names/{key}", namesHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(STATIC_DIR)))

	return httpsRedirectRouter(r)
}

func httpsRedirectRouter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proto := r.Header.Get("x-forwarded-proto")
		if strings.ToLower(proto) == "http" {
			http.Redirect(w, r, fmt.Sprintf("https://%s%s", r.Host, r.URL), http.StatusPermanentRedirect)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(go90s.GetRandomName()))
}

func namesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	w.Write([]byte(go90s.GetName(key)))
}
