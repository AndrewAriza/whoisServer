package utils

import (
	"encoding/json"
	"net/http"

	"github.com/fatih/color"
)

func Catch(w http.ResponseWriter, err error) {
	if err != nil {
		warning := color.New(color.FgRed).PrintlnFunc()
		payload := map[string]string{"message": err.Error()}
		response, _ := json.Marshal(payload)

		warning("ERROR :::", payload)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(response)
		panic(nil)
	}
}
