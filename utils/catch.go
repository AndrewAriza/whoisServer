package utils

import (
	"encoding/json"
	"net/http"

	"github.com/fatih/color"
)

func Catch(w http.ResponseWriter, err error) {
	if err != nil {
		Response(w, 500, err.Error())
	}
}

func Response(w http.ResponseWriter, code int, msg string) {
	payload := map[string]string{"message": msg}
	response, _ := json.Marshal(payload)

	res := color.New(color.FgBlue).PrintlnFunc()
	res(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
