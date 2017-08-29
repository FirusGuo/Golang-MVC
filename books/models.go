package books

import (
	"fmt"
	"net/http"

	"github.com/FirusGuo/Golang-MVC/config"
	"github.com/julienschmidt/httprouter"
)

//Book struct
type Book struct {
	Color string
}

// Index  Handler: Index, Route: /index
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// books := make([]Book, 0)
	books := Book{"Yellow"}

	if err := config.TPL.ExecuteTemplate(w, "index.gohtml", books); err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println("Error:500:Cannot join template:index.gohtml into fetched data: Index, Handler: Index, Route: /Index")
		fmt.Fprintln(w, "Error:500:Cannot join template:index.gohtml into fetched data: Index, Handler: Index, Route: /Index")
		return
	}
	fmt.Println("Success:Can reach to end of handler, Handler: Index, Route: /Index")
}
