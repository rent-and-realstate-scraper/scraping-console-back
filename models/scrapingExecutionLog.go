package models

import (
	"fmt"
	"time"
)

type ScrapingExecutionLog struct {
	ScrapingID  string    `json:"scraping_id"`
	LastPiece   string    `json:"last_piece"`
	LastResult  string    `json:"last_result"`
	DateScraped time.Time `json:"date_scraped"`
	AppId       string    `json:"app_id"`
	DeviceId    string    `json:"device_id"`
}

func GetLastPiece(scrapingId string) {
	sql := fmt.Sprintf("select * from scraping_execution_log where scraping_id = %s", scrapingId)
	fmt.Println(sql)
}

func GetScrapingExecutionLog(limit int, offset int, order string) []ScrapingExecutionLog {
	if order == "" {
		order = "desc"
	}

	sql := fmt.Sprintf("select t.*, r.date_scraped, r.app_id, r.device_id from scraping_execution_log t, scraping_results r where t.last_piece = r.piece_id order by r.date_scraped %s limit %d offset %d ;", order, limit, offset)
	fmt.Println(sql)

	db := GetDb()
	results, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}

	var scrapingExecutionLogs []ScrapingExecutionLog
	for results.Next() {
		var item ScrapingExecutionLog
		err = results.Scan(&item.ScrapingID, &item.LastPiece, &item.LastResult, &item.DateScraped, &item.AppId, &item.DeviceId)
		if err != nil {
			fmt.Println(err)
		}
		scrapingExecutionLogs = append(scrapingExecutionLogs, item)
	}

	return scrapingExecutionLogs
}
