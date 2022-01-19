package data

import (
	database "bitbucket.org/traveliodev/stockbit/2/db"
)

type MysqlMoviesData struct {
}

func (d MysqlMoviesData) InsertLog(payload LogInsert) error {

	db := database.GetDB()

	result := db.Table("log").Select("log_id", "url", "action", "http_method", "request", "response", "response_time", "timestamp").Create(&payload)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
