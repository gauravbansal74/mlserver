package watcher

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/gauravbansal74/mlserver/server/model/visits"

	"gopkg.in/mgo.v2/bson"

	"github.com/gauravbansal74/mlserver/pkg/logger"
	"github.com/gauravbansal74/mlserver/server/model/dataSource"
)

// Create message
func Create(fileName string, message, status string) error {
	ds := &dataSource.Entity{
		ID:        bson.NewObjectId(),
		Name:      fileName,
		Message:   message,
		Status:    status,
		CreatedAt: time.Now(),
	}
	err := ds.Create()
	return err
}

// FileDataValidation - file validation data
func FileDataValidation(name string) (int, error) {
	file, err := os.Open(name)
	if err != nil {
		logger.Error(err, "error while reading file", logger.Fields{"fileName": name})
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = '|'
	lineCount := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			logger.Error(err, "error while reading records from CSV file", logger.Fields{
				"fileName": name,
				"line":     lineCount,
			})
			return lineCount, err
		}
		if lineCount > 0 {
			if len(record) == 3 {
				recoredDate, err := time.Parse("2006-01-02", record[0])
				if err != nil {
					logger.Error(err, "error while reading date from CSV file", logger.Fields{
						"fileName": name,
						"line":     lineCount,
					})
					return lineCount, err
				}
				total, err := strconv.ParseInt(record[2], 10, 64)
				if err != nil {
					logger.Error(err, "error while reading total visit from CSV file", logger.Fields{
						"fileName": name,
						"line":     lineCount,
					})
					return lineCount, err
				}
				newVisit := visits.Entity{
					ID:      bson.NewObjectId(),
					Date:    recoredDate,
					Website: record[1],
					Total:   total,
				}
				err = newVisit.Create()
				if err != nil {
					logger.Error(err, "error while creating records in db", logger.Fields{
						"fileName": name,
						"line":     lineCount,
					})
					return lineCount, err
				}
			} else {
				return lineCount, fmt.Errorf("invalid line data")
			}
		}
		lineCount++
	}
	return lineCount, nil
}
