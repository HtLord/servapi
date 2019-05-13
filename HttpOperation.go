package servapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/HtLord/servmodel"
)

// Can be added into the end of http handler for write json into response
// Response header will be note as (Content-Type: application/json)
// This func will return error when ResponseWriter fire it. Otherwise,
// return nil
func JsonResponseWrapper(w http.ResponseWriter, v interface{}) error {
	j, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(j)
	if err != nil {
		return err
	}

	return nil
}

// The variable is a counter for counting coming request.
// You can add CountReq in your http func handler. While handler
// catch a request it will counting it. Caution all func will assign
// the same counter.
var reqc int = 1

func CountReq() {
	fmt.Printf("---Request counting(%d)---\n", reqc)
	reqc++
}

// Major ability is dump form and body from request to visualize those data
func DumpReq(r *http.Request) {
	fmt.Printf("Request from [%s] through method[%s]\n", r.URL.Path, r.Method)
	fmt.Printf("1. Body {%v}\n", r.Body)
	fmt.Printf("2. Form {%v}\n", r.Form)
	for k, v := range r.Form {
		fmt.Printf("K:{%s} V:{%s}\n", k, v)
	}
	fmt.Printf("3. PostForm {%v}\n", r.PostForm)
}

// Reduce the code while something wrong and want to fire a redirection response or
// some sort.
func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		fmt.Println("error msg:" + err.Error())
		http.NotFound(w, r)
		return true
	}
	return false
}

func ErrorJsonResponsHandler(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		e := servmodel.ErrorResponseWrapper(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		JsonResponseWrapper(w, e)
		return true
	}
	return false
}
