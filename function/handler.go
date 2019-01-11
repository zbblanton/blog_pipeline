package function

import (
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if string(path[len(path)-1]) == "/" {
		path = path + "index.html"
	}
	http.ServeFile(w, r, "public/"+path)
}
