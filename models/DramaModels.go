package models

import (
	"go-postgres/config"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Drama struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Director     string `json:"director"`
	Release_date string `json:"release_date"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FetchAllDrama() ([]Drama, error) {
	var dramas []Drama

	conn := config.CreateConnection()

	sqlStatment := `SELECT * FROM drama`

	rows, err := conn.Query(sqlStatment)
	// defer rows.Close()

	if err != nil {
		log.Fatalf("Tidak bisa megeksekusi query %v", err)
	}

	for rows.Next() {
		var drama Drama

		err = rows.Scan(&drama.ID, &drama.Title, &drama.Director, &drama.Release_date)

		if err != nil {
			log.Fatalf("Tidak bisa mengambil data %v", err)
		}

		dramas = append(dramas, drama)
	}

	return dramas, err
}

func InsertDataDrama(title string, director string, release_date string) (Response, error) {
	var res Response

	// koneksi db
	conn := config.CreateConnection()
	defer conn.Close()

	sqlStatment := `INSERT INTO drama (title, director, release_date) VALUES ($1,$2,$3)RETURNING id`

	statment, err := conn.Prepare(sqlStatment)
	if err != nil {
		return res, err
	}

	result, err := statment.Exec(title, director, release_date)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}
	return res, nil
}

func UpdateDrama(id int, title string, director string, release_date string) (Response, error) {
	var res Response

	conn := config.CreateConnection()

	sqlStatment := `UPDATE drama SET title = $1, director = $2, release_date = $3 WHERE id = $4`

	statment, err := conn.Prepare(sqlStatment)
	if err != nil {
		return res, err
	}

	result, err := statment.Exec(title, director, release_date, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sucess"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteDrama(id int) (Response, error) {
	var res Response

	conn := config.CreateConnection()

	sqlStatment := `DELETE FROM drama WHERE id = $1`

	statement, err := conn.Prepare(sqlStatment)
	if err != nil {
		return res, err
	}

	result, err := statement.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
