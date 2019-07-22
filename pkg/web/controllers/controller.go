package controllers

import (
    "fmt"
	"net/http"
	"html/template"
	"os"
	"path"
	"github.com/sc-app/pkg/storing"
	"github.com/sc-app/pkg/handles"
	"github.com/sc-app/blockchain"
)

type Application struct {
	UserStorage	*storing.UserStorage
	HandleSession *handles.Session
	Fabric        *blockchain.FabricSetup
}

func renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data interface {}){
	lp := path.Join("pkg", "web", "templates", "layout.html")
	fp := path.Join("pkg", "web", "templates", templateName)
			
	info, err := os.Stat(lp)
	if err != nil {
		if os.IsNotExist(err) {
				http.NotFound(w, r)
		}
	}
	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)	
	}	

	info, err = os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
				http.NotFound(w, r)
		}
	}
	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)	
	}	
	
	templates, err := template.ParseFiles(lp, fp)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 Internal Server Error", 500)
	}
	templates.ExecuteTemplate(w, "layout", data)
}