package main

import (
	"net/http"
	// "encoding/json"
	"log"
	"fmt"
	"io/ioutil"
	"bytes"
	"net/http/httputil"
	"strings"
)

func GetTest(url string) []byte {
    var body []byte
    var response *http.Response
    var request *http.Request

    // tr := &http.Transport{
    //     TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
    // }

    request, err := http.NewRequest("GET", url, nil)
    if err == nil {
        request.Header.Add("Content-Type", "application/json")
        debug(httputil.DumpRequestOut(request, true))
        response, err = (&http.Client{}).Do(request)
    }

    if err == nil {
        defer response.Body.Close()
        debug(httputil.DumpResponse(response, true))
        body, err = ioutil.ReadAll(response.Body)
    }

    if err != nil {
        log.Fatalf("ERROR: %s", err)
    }

    return body
}

func PostTest(url string, data []byte) []byte {
	var body []byte
	var response *http.Response
	var request *http.Request
 
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	// }
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err == nil {
		request.Header.Add("Content-Type", "application/json")
		debug(httputil.DumpRequestOut(request, true))
		response, err = (&http.Client{}).Do(request)
	}
 
	if err == nil {
		defer response.Body.Close()
		debug(httputil.DumpResponse(response, true))
		body, err = ioutil.ReadAll(response.Body)
	}
 
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
 
	return body
}

func debug(data []byte, err error) {
    if err == nil {
        fmt.Printf("%s\n\n", data)
    } else {
        log.Fatalf("%s\n\n", err)
    }
}

func main() {

/* ADD & GET users - Init testing */
// PostTest("http://127.0.0.1:8000/addUser", []byte(body))
// fmt.Println(GetTest("http://127.0.0.1:8000/getUser/12345"))

/* ADD multiple users and LIST them - Init testing */


	// url := "http://127.0.0.1:8000/addUser"
	// userList := []string {"Ted", "Ed", "Mel", "Bob", "Jay", "Al", "Jen"}
	
	// for _, v := range userList {
	// 	body := fmt.Sprintf(`{"id": %q, "data": "APItest"}`, v)
	// 	PostTest(url, []byte(body))
	// }
	
	// GetTest("http://127.0.0.1:8000/listUser")

/* CONSENT - Init testing */

	url := "http://127.0.0.1:8080/addUser"
	idList := []string{
		"Ken",
		"Rick",
		"Morty",
		"Ray",
		"Mel",
		"Al",
		"Ted",
	}
	var body []string
	for _, v := range idList {
		body = append(body, fmt.Sprintf(`{"id": %q}`, v))
	}

	temp := strings.Join(body,",")
	temp = fmt.Sprintf(`[%s]`, temp)
	fmt.Println(body)

	PostTest(url, []byte(temp))
	url = "http://127.0.0.1:8080/getUser/Mob"

	GetTest(url)
}