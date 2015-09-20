package routes

import (
	"fmt"
	"github.com/MarvinMenzerath/UpAndRunning2/lib"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func IndexIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := SiteData{lib.GetConfiguration().Dynamic.Title}
	t, err := template.ParseFiles("views/index.html", "views/partials/styles.html", "views/partials/footer.html", "views/partials/scripts.html")

	if t != nil {
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, data)
	} else {
		fmt.Println("Error while parsing Template: ", err)
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
	}
}

func IndexStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
