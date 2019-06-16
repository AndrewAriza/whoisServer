package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/andrewariza/whoisServer/models"

	"github.com/andrewariza/whoisServer/utils"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("postgres", "postgresql://truora:cockroach@localhost:26257/cluster?sslmode=verify-full&sslrootcert=certs/ca.crt")
}

func Create(w http.ResponseWriter, domainName string, sslgrade string, domain string) {
	query, err := db.Prepare("INSERT INTO server (domain, sslgrade, ssllab, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW());")
	utils.Catch(w, err)

	_, err = query.Exec(domainName, sslgrade, domain)
	utils.Catch(w, err)
	defer query.Close()
}

func Find(w http.ResponseWriter, domainName string) (string, models.Domain) {
	query, err := db.Prepare(`SELECT sslgrade, ssllab FROM server WHERE ssllab @> '{"host": "truora.com"}' AND created_at < NOW() - INTERVAL '1 hour' ORDER BY created_at DESC LIMIT 1;`)
	utils.Catch(w, err)

	var (
		sslgrade = "M"
		data     string
	)
	err = query.QueryRow().Scan(&sslgrade, &data)

	domain := models.Domain{}
	if err == nil {
		err = json.Unmarshal([]byte(data), &domain)
		utils.Catch(w, err)
	}

	defer query.Close()
	return sslgrade, domain
}

func Record(w http.ResponseWriter, r *http.Request) {
	query, err := db.Prepare(`SELECT DISTINCT domain, sslgrade FROM server LIMIT 10;`)
	utils.Catch(w, err)

	record := models.Record{}
	rows, err := query.Query()
	utils.Catch(w, err)

	for rows.Next() {
		item := models.Item{}
		err := rows.Scan(&item.Domain, &item.Sslgrade)
		utils.Catch(w, err)
		record.Items = append(record.Items, item)
	}

	response, _ := json.Marshal(record)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

	defer rows.Close()
}
