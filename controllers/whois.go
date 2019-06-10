package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/andrewariza/whoisServer/models"
	"github.com/andrewariza/whoisServer/utils"
	"github.com/go-chi/chi"
	"github.com/gocolly/colly"
)

func Whois(w http.ResponseWriter, r *http.Request) {
	domainName := chi.URLParam(r, "domain")

	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domainName)
	utils.Catch(w, err)

	var domain models.Domain
	err = json.NewDecoder(response.Body).Decode(&domain)
	utils.Catch(w, err)

	cluster := models.Cluster{}

	cluster.Servers = []models.Server{}
	for _, element := range domain.Endpoints {
		ip := Ip(element.IpAddress, w)
		server := models.Server{}

		server.Address = element.IpAddress
		if server.SslGrade == "" {
			server.SslGrade = element.StatusMessage
		}
		server.Country = ip.CountryCode
		server.Owner = ip.Isp
		cluster.Servers = append(cluster.Servers, server)
	}

	var (
		logo  string
		title string
	)

	c := colly.NewCollector()

	c.OnHTML("head", func(e *colly.HTMLElement) {
		logo = e.ChildAttr("link[rel='shortcut icon']", "href")
		title = e.ChildText("title")
		if logo == "" {
			logo = e.ChildAttr("link[rel='icon']", "href")
		}
	})

	c.Visit("https://" + domainName)

	cluster.ServersChanged = true
	cluster.SslGrade = "B"
	cluster.PreviousSslGrade = "A+"
	cluster.Logo = logo
	cluster.Title = title

	if domain.Status == "ERROR" {
		cluster.IsDown = true
	} else {
		cluster.IsDown = false
	}

	res, err := json.Marshal(cluster)
	utils.Catch(w, err)

	w.Write([]byte(res))
}
