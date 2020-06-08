package database

var (
	CreateEventTable = "CREATE TABLE IF NOT EXISTS Event(" +
		"id serial primary key, " +
		"tag varchar (50), " +
		"name varchar (150), " +
		"available integer, " +
		"capacity integer, " +
		"status integer, " +
		"frontends text, " +
		"exercises text, " +
		"started_at varchar (100), " +
		"finish_expected varchar (100), " +
		"finished_at varchar (100));"

	CreateTeamsTable = "CREATE TABLE IF NOT EXISTS Team(" +
		"id serial primary key, " +
		"tag varchar (50), " +
		"event_id integer, " +
		"email varchar (50), " +
		"name varchar (50), " +
		"password varchar (250), " +
		"created_at varchar (100), " +
		"last_access varchar (100), " +
		"solved_challenges text);"

	AddTeamQuery = "INSERT INTO team (tag, event_id, email, name, password, created_at, last_access, solved_challenges)" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	AddEventQuery = "INSERT INTO event (tag, name, available, capacity, frontends, status, exercises, started_at, finish_expected)" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	UpdateEventFinishDate       = "UPDATE event SET finished_at = $2 WHERE tag = $1"
	UpdateEventStatus           = "UPDATE event SET status = $2 WHERE tag = $1 "
	UpdateEventLastaccessedDate = "UPDATE team SET last_access = $2 WHERE tag = $1"
	UpdateTeamSolvedChl         = "UPDATE team SET solved_challenges = $2 WHERE tag = $1"

	QuerySolvedChls = "SELECT solved_challenges FROM team WHERE tag=$1"
	QueryEventTable = "SELECT * FROM event"

	QueryEventId    = "SELECT id FROM event WHERE tag=$1 and finished_at is null"
	QueryEventTeams = "SELECT * FROM team WHERE event_id=$1"

	QueryEventStatus = "SELECT status FROM event WHERE tag=$1"
)
