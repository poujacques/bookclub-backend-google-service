package googleBooks

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
)

const BASE_URI = "https://www.googleapis.com/books/v1"

func GetVolumesResult(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	resp, err := http.Get(fmt.Sprintf("%s/volumes?q=%s", BASE_URI, url.QueryEscape(q)))
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading http response: %s\n", err)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}

func GetVolumeByID(w http.ResponseWriter, r *http.Request) {
	volumeID := chi.URLParam(r, "volumeID")
	resp, err := http.Get(fmt.Sprintf("%s/volumes/%s", BASE_URI, volumeID))
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading http response: %s\n", err)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}
