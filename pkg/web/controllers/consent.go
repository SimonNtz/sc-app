package controllers

import (
    "fmt"
	"net/http"
	// "os"
	"encoding/json"
)

func (app *Application) ConsentHandler(w http.ResponseWriter, r *http.Request) {
	/* CHECK IF STORAGE HAS BEEN INITIATED */
	var tmpl string
	session, _ := store.Get(r, "session-test")
	/* centralize for the package*/
	if usrId := session.Values["PID"]; usrId != nil {
		// tst, _ := json.Marshal(contact)
		r.ParseForm()
		// var consents map[int]interface{}
		// consentsJSON, _ := json.Marshal(r.Form)

		// fmt.Println(string(consentsJSON))
		app.UserStorage.AddConsent(usrId.(string), r.Form)
		app.UserStorage.Print()

		// app.UserStorage.AddSignature(usrId.(string), signature)
		
		// IF USER'S PROFIL SIGNED

		//PUBLISH INTO BLOCKCHAIN via GO ROUTINE

		usr := app.UserStorage.Get(usrId.(string))

		// h := json.Marshall(fmt.Sprintf(`{%+v}`, usr.Consents))
	
		c := struct {
			Consents 	*map[string][]string	 `json:"consents"`
			Signature	string           		`json:"signature"`
		}{Consents: &usr.Consents, Signature: usr.Signature}
	
		b, err := json.MarshalIndent(&c, "", "\t")
		if err != nil {
			fmt.Println("error:", err)
		}
		// os.Stdout.Write(b)
	
		_ = b // REMOVE
		response, err := app.Fabric.InvokeHello(usr.Id, b)
		if err != nil {
			fmt.Printf("Unable to writte consent into chaincode: %v\n", err)
		} else {
			fmt.Printf("Response from the Smart-Contract hello: %s\n", response)
		}
		tmpl = "consent.html"
	} else {
		tmpl = "unsuscribed.html"
	}
	/*Send transaction informations as dataTemplate*/
	renderTemplate(w, r, tmpl, nil)
}