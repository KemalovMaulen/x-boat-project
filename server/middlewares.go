package server

import (
	"net/http"
	"encoding/json"
	"bytes"
	"github.com/salambayev/x-boat-project/logger"
	"time"
	"io/ioutil"
	"strings"
	"fmt"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

type HttpEndpoint func(w http.ResponseWriter, r *http.Request) HttpResponse

type HttpResponse interface {
	Headers() map[string]string
	Response() interface{}
	StatusCode() int
}

type Response struct {
	Status     int
	Data       interface{}
	HeaderData map[string]string
}

func (e *Response) Response() interface{} {
	return e.Data
}

func (e *Response) StatusCode() int {
	return e.Status
}

func (e *Response) Headers() map[string]string {
	return e.HeaderData
}

func Json(fn HttpEndpoint) HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		d := fn(w, r)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		for k, v := range d.Headers() {
			w.Header().Set(k, v)
		}

		w.WriteHeader(d.StatusCode())
		json.NewEncoder(w).Encode(d.Response())
	}
}

func Logging(log logger.Logger, fn HttpEndpoint) HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) HttpResponse {
		start := time.Now()

		// Read the content
		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(r.Body)
		}
		// Restore the io.ReadCloser to its original state
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// Use the content
		bodyString := string(bodyBytes)

		d := fn(w, r)
		dBytes, _ := json.Marshal(d.Response())

		if d.StatusCode() > 299 || d.StatusCode() < 200 {
			log.Warn(LogRequest(r), " body= ", bodyString, time.Since(start), " ", d.StatusCode(), " ", string(dBytes))
		} else {
			log.Debug(LogRequest(r), " body= ", bodyString, time.Since(start), " ", d.StatusCode(), " ", string(dBytes))
		}

		return d
	}
}

func LogRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	urlPath := fmt.Sprintf("%v %v", r.Method, r.URL)
	request = append(request, urlPath)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, " ")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, " ") + " "
}


func OK(d interface{}) *Response {
	resp := &Response{}
	resp.Status = http.StatusOK
	resp.Data = d
	return resp
}

func Created(d interface{}) *Response {
	resp := &Response{}
	resp.Status = http.StatusCreated
	resp.Data = d
	return resp
}