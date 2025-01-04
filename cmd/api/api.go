package api

import "database/sql"

type APIServer struct {
	adress string
	db     *sql.DB
}

func NewAPIServer(adress string, db *sql.DB) *APIServer {
	return &APIServer{adress: adress, db: db}
}
