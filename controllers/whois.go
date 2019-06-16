package controllers

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/andrewariza/whoisServer/models"
	"github.com/andrewariza/whoisServer/utils"
	"github.com/go-chi/chi"
	"github.com/gocolly/colly"
)

func Whois(w http.ResponseWriter, r *http.Request) {
	domainName := strings.ToLower(chi.URLParam(r, "domain"))

	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domainName)
	utils.Catch(w, err)

	domain := models.Domain{}
	err = json.NewDecoder(response.Body).Decode(&domain)
	utils.Catch(w, err)

	server, err := json.Marshal(domain)
	utils.Catch(w, err)

	ssl := []string{
		"M", "T",
		"F-", "F", "F+",
		"E-", "E", "E+",
		"D-", "D", "D+",
		"C-", "C", "C+",
		"B-", "B", "B+",
		"A-", "A", "A+"}

	grade := 0

	cluster := models.Cluster{}

	cluster.Servers = []models.Server{}
	for _, element := range domain.Endpoints {
		ip := Ip(element.IpAddress, w)
		server := models.Server{}

		server.Address = element.IpAddress
		if element.Grade == "" {
			server.SslGrade = element.StatusMessage
		} else {
			server.SslGrade = element.Grade
			for i, _ := range ssl {
				if (ssl[i] == server.SslGrade) && ((i < grade) || (grade == 0)) {
					grade = i
				}
			}
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
		logo = e.ChildAttr(`link[rel="shortcut icon"]`, "href")
		title = e.ChildText("title")
		if logo == "" {
			logo = e.ChildAttr(`link[rel="icon"]`, "href")
		}
	})

	c.Visit("http://" + domainName)

	sslGrade := ssl[grade]

	previousSslGrade, previousDomain := Find(w, domainName)

	cluster.ServersChanged = !reflect.DeepEqual(previousDomain, domain)
	cluster.SslGrade = ssl[grade]
	cluster.PreviousSslGrade = previousSslGrade
	cluster.Logo = logo
	cluster.Title = title

	if domain.Status == "ERROR" {
		cluster.IsDown = true
	} else {
		cluster.IsDown = false
	}

	Create(w, domainName, sslGrade, string(server))

	res, err := json.Marshal(cluster)
	utils.Catch(w, err)

	w.Write(res)
}
