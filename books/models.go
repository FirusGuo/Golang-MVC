import books

import (
	"net/http"

  "github.com/FirusGuo/Golang-MVC/config"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

// Index  Handler: Index, Route: /index
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

if err := config.TPL.ExecuteTemplate(w, "index.gohtml", users); err != nil {
  http.Error(w, http.StatusText(500), 500)
  fmt.Println("Error:500:Cannot join template:index.gohtml into fetched data: Index, Handler: Index, Route: /Index")
  fmt.Fprintln(w, "Error:500:Cannot join template:index.gohtml into fetched data: Index, Handler: Index, Route: /Index")
  return
}
fmt.Println("Success:Can reach to end of handler, Handler: Index, Route: /Index")
}
