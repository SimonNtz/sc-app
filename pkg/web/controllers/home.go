package controllers

import (
    "fmt"
	"net/http"
	"github.com/gorilla/mux"
	// "time"
	"github.com/gorilla/sessions"


	// "github.com/gorilla/session"
)

/* Export to a model in subpackage */

type Affiliations struct {
	Names []string
}

var (
	store = sessions.NewCookieStore([]byte("PID"))
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	
	// if r.Method != http.MethodPost {
	// 	fmt.Printf("Wrong method in App")
	// 	json.NewEncoder(w).Encode([]byte("Wrong Method in App"))
	// 	return
	// }

/* Check form's value completness */
	var tmpl string
	params := mux.Vars(r)

	if usr:= app.UserStorage.Get(params["id"]); usr.Id != "" {
		fmt.Printf("COOKIE BAKED -> user %q connected \n", usr.Id)
		session, _ := store.Get(r, "session-test")
		session.Values["PID"] = usr.Id
		session.Save(r, w)
		tmpl = "home.html"
	} else {	
		tmpl = "unsuscribed.html"
	}

	dataTemplate := Affiliations {
			Names: []string{"ITU", "ONU", "OMC"},
	}
	
	renderTemplate(w, r, tmpl, dataTemplate)
}
