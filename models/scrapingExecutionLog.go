package models

import "fmt"

type ScrapingExecutionLog struct {
	ScrapingId string `json:"scraping_id"`
	LastPiece  string `json:"last_piece"`
	LastResult string `json:"last_result"`
}

func GetLastPiece(scrapingId string) {
	sql := fmt.Sprintf("select * from scraping_execution_log where scraping_id = %s", scrapingId)
	fmt.Println(sql)
}
