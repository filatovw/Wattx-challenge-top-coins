package api

import "net/http"

func WriteError(w http.ResponseWriter, header http.Header, err error) {
	sterr := ToStatusError(err)
	w.Write([]byte(sterr.message))
}

func WriteOK(w http.ResponseWriter, header http.Header, body []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
