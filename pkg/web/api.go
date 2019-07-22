package web

import (
    "fmt"
    // "encoding/json"
    "github.com/gorilla/mux"
    "log"
	"net/http"
	// "io/ioutil"
	// "github.com/sc-app/pkg/users"
	// "github.com/sc-app/pkg/storing"
	"github.com/sc-app/pkg/web/controllers"
	// "github.com/sc-app/pkg/storing"
	// "html/template"
	"time"
	"os"
	// "path"
)

/* Flexible design for POST call rules are: 
	- AddUser Id mandat., info and not mandat.
	- AddUserInfo Id mandat., info, subject
*/

// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	/* CHECK IF STORAGE HAS BEEN INITIATED */
// 	params := mux.Vars(r)
// 	if  userId := params["id"]; userStorage.Get(userId).Id != ""   {
// 		fmt.Printf("Found")
// 	}

// 	u:= userStorage.Get(params["id"])
// 	json.NewEncoder(w).Encode(u)
// } 

// func ListUser(w http.ResponseWriter, r *http.Request) {
// 	/* CHECK IF STORAGE HAS BEEN INITIATED */
// 	var userList []string 

// 	for _, i := range userStorage.UserRepo {
// 		userList = append(userList, i.Id)
// 	}	
	
// 	json.NewEncoder(w).Encode(userList)
// } 



// func main() {
func Serve(app *controllers.Application) {
	router := mux.NewRouter()
	router.HandleFunc("/app", app.AppHandler).Methods("POST")
	router.HandleFunc("/consent", app.ConsentHandler).Methods("POST")
	router.PathPrefix("/pkg/web/assets/").Handler(http.FileServer(http.Dir(".")))
	router.HandleFunc("/home.html/{id}", app.HomeHandler).Methods("GET")
	// router.Handle("/home.html/{id}", login(controllers.HomeHandler)).Methods("GET")

	// router.HandleFunc("/listConsent", ListConsent).Methods("GET")
	// router.HandleFunc("/listUser", ListUser).Methods("GET")
	// router.HandleFunc("/getUser/{id}", GetUser).Methods("GET")
	// router.HandleFunc("/addUser", AddUsers).Methods("POST")
	srv := &http.Server{
        Handler:      router,
        Addr:         "127.0.0.1:8080",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
	log.Fatal(srv.ListenAndServe())
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "/home.html", http.StatusTemporaryRedirect)
	// })

}


func login(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(mux.Vars(r)["id"])
		f(w, r)
	}
}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
	port = "9090"
	fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}