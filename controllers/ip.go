package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/andrewariza/whoisServer/models"
	"github.com/andrewariza/whoisServer/utils"
)

func Ip(endpoint string, w http.ResponseWriter) models.Ip {
	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Get("http://ip-api.com/json/" + endpoint)
	utils.Catch(w, err)

	ip := models.Ip{}
	err = json.NewDecoder(response.Body).Decode(&ip)
	utils.Catch(w, err)

	return ip
}
