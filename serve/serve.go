package serve

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Serve() {
	route := mux.NewRouter()
	route.HandleFunc("/health_check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.Handle("/", route)
	err := http.ListenAndServe(":8088", nil)
	fmt.Printf("err=%v", err)
}
