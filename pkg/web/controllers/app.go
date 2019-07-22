package controllers

import (
    "fmt"
	"net/http"
	"encoding/json"
	// "github.com/Gohandler"
)

/* Export to a model in subpackage */

type ContactInfos struct {
	Fname 			string		`json:"firstname"` 
	Lname 			string		`json:"lastname"`
	Email   		string		`json:"email"`
	Afln		 	string		`json:"affiliation"`
}


func (app *Application) AppHandler(w http.ResponseWriter, r *http.Request) {
	// CHECK IF FORM IS COMPLETE
	var tmpl string
	r.ParseForm()
	contact := ContactInfos {
		Fname: r.FormValue("fname"),   
		Lname: r.FormValue("lname"),   
		Email: r.FormValue("email"),   
		Afln : r.FormValue("afln"),	
	}

	/* Technically auth to LHS is not needed however */
	// 1-AUTH
	// 2-POST INFO
	// 3-STORE RESULTING SIGNATURE

	/*Do something with contact*/

	session, _ := store.Get(r, "session-test")
	/* centralize for the package*/
	if usrId := session.Values["PID"]; usrId != nil {
		tst, _ := json.Marshal(contact)
		hdl := map[string]string{ 
			"index": "120",  
			"type": "HS_USER",
			"value": string(tst),
		}
		/* Remove hardcoded values */
		app.HandleSession.PostHandle(fmt.Sprintf("11.test/smart-consent/run0/users/%s", usrId), hdl)
		signature := app.HandleSession.SignHandle(fmt.Sprintf("11.test/smart-consent/run0/users/%s", usrId))
		app.UserStorage.AddSignature(usrId.(string), signature)

		tmpl = "app.html"

	} else {
		tmpl = "unsuscribed.html"
	}
	
	renderTemplate(w, r, tmpl, app.UserStorage.Topics)
}
