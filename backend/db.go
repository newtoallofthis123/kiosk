package main

import (
	"database/sql"
	"strings"

	_ "github.com/lib/pq"
	"github.com/newtoallofthis123/ranhash"
)

type DbInstance struct {
	Db *sql.DB
}

func NewDb(env *Env) (*DbInstance, error) {
	db, err := sql.Open("postgres", GetDbConnString(env))
	if err != nil {
		return nil, err
	}

	dbInstance := &DbInstance{Db: db}
	err = dbInstance.initDb()
	if err != nil {
		return nil, err
	}

	return dbInstance, nil
}

func (pq *DbInstance) initDb() error {
	query := `
	CREATE TABLE IF NOT EXISTS records(
		id VARCHAR(8) PRIMARY KEY,
		name TEXT NOT NULL,
		problems TEXT NOT NULL,
		diagnosis TEXT NOT NULL,
		treatment TEXT NOT NULL,
		review TEXT NOT NULL,
		recommendations TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW()
	);
	`

	_, err := pq.Db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

type Record struct {
	Id              string   `json:"id"`
	Name            string   `json:"name"`
	Problems        []string `json:"problems"`
	Diagnosis       []string `json:"diagnosis"`
	Treatment       []string `json:"treatment"`
	Review          string   `json:"review"`
	Recommendations []string `json:"recommendations"`
	CreatedAt       string   `json:"created_at"`
}

type ConstructRequest struct {
	Name            string   `json:"name"`
	Problems        []string `json:"problems"`
	Diagnosis       []string `json:"diagnosis"`
	Treatment       []string `json:"treatment"`
	Review          string   `json:"review"`
	Recommendations []string `json:"recommendations"`
}

func (pq *DbInstance) CreateRecord(req *ConstructRequest) (string, error) {
	id := ranhash.GenerateRandomString(8)
	query := `
	INSERT INTO records (id, name, problems, diagnosis, treatment, review, recommendations)
	VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	_, err := pq.Db.Exec(query, id, req.Name, strings.Join(req.Problems, "|"), strings.Join(req.Diagnosis, "|"), strings.Join(req.Treatment, "|"), req.Review, strings.Join(req.Recommendations, "|"))
	if err != nil {
		return "", err
	}

	return id, nil
}

func (pq *DbInstance) GetRecord(id string) (*Record, error) {
	query := `
	SELECT id, name, problems, diagnosis, treatment, review, recommendations, created_at
	FROM records
	WHERE id = $1;
	`

	row := pq.Db.QueryRow(query, id)
	problems := ""
	diagnosis := ""
	treatment := ""
	recommendations := ""
	record := &Record{}
	err := row.Scan(&record.Id, &record.Name, &problems, &diagnosis, &treatment, &record.Review, &recommendations, &record.CreatedAt)
	if err != nil {
		return nil, err
	}

	record.Problems = strings.Split(problems, "|")
	record.Diagnosis = strings.Split(diagnosis, "|")
	record.Treatment = strings.Split(treatment, "|")
	record.Recommendations = strings.Split(recommendations, "|")

	return record, nil
}

func (pq *DbInstance) DeleteRecord(id string) error {
	query := `
	DELETE FROM records
	WHERE id = $1;
	`

	_, err := pq.Db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (pq *DbInstance) UpdateRecord(id string, req *Record) error {
	query := `
	UPDATE records
	SET name = $1, problems = $2, diagnosis = $3, treatment = $4, review = $5, recommendations = $6
	WHERE id = $7;
	`

	_, err := pq.Db.Exec(query, req.Name, strings.Join(req.Problems, "|"), strings.Join(req.Diagnosis, "|"), strings.Join(req.Treatment, "|"), req.Review, strings.Join(req.Recommendations, "|"), id)
	if err != nil {
		return err
	}

	return nil
}

func (pq *DbInstance) ListRecords() ([]*Record, error) {
	query := `
	SELECT id FROM records;
	`

	rows, err := pq.Db.Query(query)
	if err != nil {
		return nil, err
	}

	records := []*Record{}
	for rows.Next() {
		id := ""
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		record, err := pq.GetRecord(id)
		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}
