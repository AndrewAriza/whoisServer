package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/andrewariza/whoisServer/utils"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgresql://truora:cockroach@localhost:26257/cluster?sslmode=verify-full&sslrootcert=certs/ca.crt")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
}

func Create(w http.ResponseWriter, domainName string, domain string) {
	if _, err := db.Exec(fmt.Sprintf(`INSERT INTO server (domain, ssllab, created_at, updated_at) 
	VALUES ('%s', '%s', NOW(), NOW());`, domainName, domain)); err != nil {
		utils.Catch(w, err)
	}
}

func Find(w http.ResponseWriter, domainName string, domain string) {
	if _, err := db.Exec(fmt.Sprintf(`INSERT INTO server (domain, ssllab, created_at, updated_at) 
	VALUES ('%s', '%s', NOW(), NOW());`, domainName, domain)); err != nil {
		utils.Catch(w, err)
	}
}
