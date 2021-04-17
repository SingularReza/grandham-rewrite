package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	handler "github.com/SingularReza/grandham-rewrite/handlers"
	mux "github.com/gorilla/mux"
	cors "github.com/rs/cors"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

// export this to handlers and import here
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	router := mux.NewRouter()

	library := router.PathPrefix("/library").Subrouter()
	library.HandleFunc("/create", handler.CreateLibrary)
	library.HandleFunc("/list", handler.GetLibraryList)
	library.HandleFunc("/items", handler.GetLibraryItems)

	item := router.PathPrefix("/item").Subrouter()
	item.HandleFunc("/info", handler.GetItemInfo)

	image := router.PathPrefix("/image/")
	image.Handler(http.StripPrefix("/image/", http.FileServer(http.Dir("./images/"))))

	spa := spaHandler{staticPath: "dist", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	srv := &http.Server{
		Handler:      c.Handler(router),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  20 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
