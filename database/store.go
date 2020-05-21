package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aau-network-security/haaukins-store/model"
	pb "github.com/aau-network-security/haaukins-store/proto"
	_ "github.com/lib/pq"
	"log"
	"strings"
	"sync"
	"time"
)

const handleNullConversionError = "converting NULL to string is unsupported"

var (
	timeFormat = "2006-01-02 15:04:05"
	OK         = "ok"
)

type store struct {
	m  sync.Mutex
	db *sql.DB
}

type Store interface {
	AddEvent(*pb.AddEventRequest) (string, error)
	AddTeam(*pb.AddTeamRequest) (string, error)
	GetEvents() ([]model.Event, error)
	GetTeams(string) ([]model.Team, error)
	UpdateTeamSolvedChallenge(*pb.UpdateTeamSolvedChallengeRequest) (string, error)
	UpdateTeamLastAccess(*pb.UpdateTeamLastAccessRequest) (string, error)
	UpdateEventFinishDate(*pb.UpdateEventRequest) (string, error)
}

func NewStore(conf *model.Config) (Store, error) {
	db, err := NewDBConnection(conf)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}
	err = InitTables(db)
	if err != nil {
		log.Printf("failed to init tables: %v", err)
		return nil, err
	}
	return &store{db: db}, nil
}

func NewDBConnection(conf *model.Config) (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Pass, conf.DB.Name)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (s *store) AddEvent(in *pb.AddEventRequest) (string, error) {
	s.m.Lock()
	defer s.m.Unlock()

	_, err := s.db.Exec(ADD_EVENT_QUERY, in.Tag, in.Name, in.Available, in.Capacity, in.Frontends, in.Exercises, in.StartTime, in.ExpectedFinishTime)

	if err != nil {
		return "", err
	}
	return "Event correctly added!", nil
}

func (s *store) AddTeam(in *pb.AddTeamRequest) (string, error) {
	s.m.Lock()
	defer s.m.Unlock()

	now := time.Now()
	nowString := now.Format(timeFormat)

	var eventId int
	if err := s.db.QueryRow(QUERY_EVENT_ID, in.EventTag).Scan(&eventId); err != nil {
		return "", err
	}

	_, err := s.db.Exec(ADD_TEAM_QUERY, in.Id, eventId, in.Email, in.Name, in.Password, nowString, nowString, "[]")
	if err != nil {
		return "", err
	}
	return "Team correctly added!", nil
}

func (s *store) GetEvents() ([]model.Event, error) {

	s.m.Lock()
	defer s.m.Unlock()

	rows, err := s.db.Query(QUERY_EVENT_TABLE)
	if err != nil {
		return nil, err
	}
	var events []model.Event
	for rows.Next() {

		event := new(model.Event)
		err := rows.Scan(&event.Id, &event.Tag, &event.Name, &event.Available, &event.Capacity, &event.Frontends,
			&event.Exercises, &event.StartedAt, &event.ExpectedFinishTime, &event.FinishedAt)
		if err != nil && !strings.Contains(err.Error(), handleNullConversionError){
			return nil, err
		}
		events = append(events, *event)
	}

	return events, nil
}

func (s *store) GetTeams(tag string) ([]model.Team, error) {
	s.m.Lock()
	defer s.m.Unlock()

	var eventId int
	if err := s.db.QueryRow(QUERY_EVENT_ID, tag).Scan(&eventId); err != nil && !strings.Contains(err.Error(), "no rows in result set"){
		return nil, err
	}

	rows, err := s.db.Query(QUERY_EVENT_TEAMS, eventId)
	if err != nil {
		return nil, err
	}

	var teams []model.Team
	for rows.Next() {

		team := new(model.Team)
		err := rows.Scan(&team.Id, &team.Tag ,&team.EventId, &team.Email, &team.Name, &team.Password, &team.CreatedAt,
			&team.LastAccess, &team.SolvedChallenges)
		if err != nil && !strings.Contains(err.Error(), handleNullConversionError) {
			return nil, err
		}
		teams = append(teams, *team)
	}
	return teams, nil
}

func (s *store) UpdateTeamSolvedChallenge(in *pb.UpdateTeamSolvedChallengeRequest) (string, error) {
	s.m.Lock()
	defer s.m.Unlock()

	type Challenge struct {
		Tag         string `json:"tag"`
		CompletedAt string `json:"completed-at"`
	}

	var solvedChallenges []Challenge
	var solvedChallengesDB string

	if err := s.db.QueryRow(QUERY_SOLVED_CHLS, in.TeamId).Scan(&solvedChallengesDB); err != nil {
		return "", err
	}

	if err := json.Unmarshal([]byte(solvedChallengesDB), &solvedChallenges); err != nil {
		return "", err
	}

	for _, sc := range solvedChallenges {
		if sc.Tag == in.Tag {
			return "", errors.New("challenge already solved")
		}
	}

	solvedChallenges = append(solvedChallenges, Challenge{
		Tag:         in.Tag,
		CompletedAt: in.CompletedAt,
	})

	newSolvedChallengesDB, _ := json.Marshal(solvedChallenges)

	_, err := s.db.Exec(UPDATE_TEAM_SOLVED_CHL, in.TeamId, string(newSolvedChallengesDB))
	if err != nil {
		return "", err
	}

	return OK, nil
}

func (s *store) UpdateTeamLastAccess(in *pb.UpdateTeamLastAccessRequest) (string, error) {
	s.m.Lock()
	defer s.m.Unlock()

	_, err := s.db.Exec(UPDATE_EVENT_LASTACCESSED_DATE, in.TeamId, in.AccessAt)
	if err != nil {
		return "", err
	}

	return OK, nil
}

func (s *store) UpdateEventFinishDate(in *pb.UpdateEventRequest) (string, error) {
	s.m.Lock()
	defer s.m.Unlock()

	_, err := s.db.Exec(UPDATE_EVENT_FINISH_DATE, in.EventId, in.FinishedAt)
	if err != nil {
		return "", err
	}

	return OK, nil
}
