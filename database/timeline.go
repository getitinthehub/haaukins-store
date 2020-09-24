package database

import (
	"database/sql"
	"log"
	"math"
	"strings"
	"time"

	"github.com/aau-network-security/haaukins-store/model"
)

// calculateCost will return a map which is
// time and number of running vms in total for
// given time, it is like a timeSeries
func calculateCost(db *sql.DB) (map[string]int32, error) {

	sT, err := getEarliestDate(db)
	if err != nil {
		log.Fatalf("Error get earliest date %v", err)
	}

	fT, err := getLastDate(db)
	if err != nil {
		log.Fatalf("Error get latest date %v", err)
	}
	timeSeries := getDates(sT, fT)

	timeSeriesCount := make(map[string]int32)
	for _, time := range timeSeries {
		timeSeriesCount[time.Format(TimeFormat)] = 0
	}

	events := getEvents(db)

	for _, e := range events {
		sT, err := time.Parse(time.RFC3339, e.StartedAt)
		if err != nil {
			log.Fatalf("Error happened %v ", err)
		}
		fT, err := time.Parse(time.RFC3339, e.ExpectedFinishTime)
		if err != nil {
			log.Fatalf("Error happened %v ", err)
		}
		if isInvalidDate(sT) && isInvalidDate(fT) {
			log.Fatalf("Invalid date formatting !!! ")
		}
		invalidTime, err := time.Parse(TimeFormat, "0001-01-01 00:00:00")
		if err != nil {
			log.Fatalf("Error happened %v ", err)
		}
		iterateOver := getDates(sT, fT)
		teamCount := getTeamsCount(db, int(e.Id))
		available := int(e.Available)
		for _, f := range iterateOver {
			timeSeriesCount[f.Format(TimeFormat)] += int32(available) + int32(teamCount)
		}
		delete(timeSeriesCount, invalidTime.Format(TimeFormat))
	}

	return timeSeriesCount, nil
}

// getEvents will query event table
// without any condition
func getEvents(db *sql.DB) []model.Event {
	r, err := db.Query(QueryEventTable)
	if err != nil {
		log.Fatalf("Error on executing query %v", err)
	}
	var events []model.Event
	for r.Next() {
		event := new(model.Event)
		err := r.Scan(&event.Id, &event.Tag, &event.Name, &event.Available, &event.Capacity, &event.Status, &event.Frontends,
			&event.Exercises, &event.StartedAt, &event.ExpectedFinishTime, &event.FinishedAt, &event.CreatedBy, &event.OnlyVPN)
		if err != nil && !strings.Contains(err.Error(), "Null conversion error ") {
			log.Fatalf("Error on scanning query %v", err)
		}
		events = append(events, *event)
	}

	return events
}

// getEarliestDate returns largest (finishDate) date from events table
func getLastDate(db *sql.DB) (time.Time, error) {
	var latestFinishTime time.Time
	r, err := db.Query(LatestDate)
	if err != nil {
		log.Fatalf("Latest Date query error %v", err)
	}
	latestTime := new(time.Time)
	for r.Next() {
		if err := r.Scan(&latestTime); err != nil {
			return time.Time{}, err
		}
		latestFinishTime = *latestTime
	}
	return latestFinishTime, nil

}

// getEarliestDate returns smallest date from events table
func getEarliestDate(db *sql.DB) (time.Time, error) {
	var earliestStartTime time.Time
	r, err := db.Query(EarliestDate)
	if err != nil {
		log.Fatalf("Earliest Date query error %v", err)
	}
	earliestTime := new(time.Time)
	for r.Next() {
		if err := r.Scan(&earliestTime); err != nil {
			return time.Time{}, err
		}
		earliestStartTime = *earliestTime
	}
	return earliestStartTime, nil
}

// isValidDate function is ensuring that
// time parsing done correctly and successfully
func isInvalidDate(t time.Time) bool {
	if t == time.Date(0001, 01, 01, 00, 00, 00, 0000, time.UTC) {
		//log.Println("Error in parsing; invalid date 0001-01-01 00:00:00 +0000 UTC  ")
		return true
	}
	return false
}

// getTemasCount return number of team on given eventID
func getTeamsCount(db *sql.DB, eventID int) int {
	var count int
	if err := db.QueryRow(QueryTeamCount, eventID).Scan(&count); err != nil {
		log.Fatalf("Query row error postgres %v", err)
	}
	return count
}

// daysInDates returns number of days in given two dates
func daysInDates(sT, fT time.Time) int {
	days := fT.Sub(sT).Hours() / 24
	return int(math.Round(days))
}

// return list of dates between given two date
// will generate list of date
func getDates(sT, fT time.Time) []time.Time {
	var dates []time.Time

	// zeroing hour:minute:second and nanosecond and setting zone to UTC
	sT = zeroTime(sT)
	fT = zeroTime(fT)
	// calculate # of days in between dates
	days := daysInDates(sT, fT)
	var count int
	for count <= days {
		date := sT
		sT = date.AddDate(0, 0, 1)
		dates = append(dates, date)
		count++
	}
	return dates
}

// zeroTime will set hour minute and second to zero
// also sets time location
func zeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0000, time.UTC)
}
