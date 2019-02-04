package api

import "net/http"

func WriteError(w http.ResponseWriter, headers map[string]string, err error) {
	sterr := ToStatusError(err)
	w.WriteHeader(sterr.Status())
	for k, v := range headers {
		w.Header().Add(k, v)
	}
	w.Write([]byte(sterr.message))
}

func WriteOK(w http.ResponseWriter, headers map[string]string, body []byte) {
	w.WriteHeader(http.StatusOK)
	for k, v := range headers {
		w.Header().Add(k, v)
	}
	w.Write(body)
}
