package database

var (
	CREATE_EVENT_TABLE = "CREATE TABLE IF NOT EXISTS Event(" +
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

	CREATE_TEAMS_TABLE = "CREATE TABLE IF NOT EXISTS Team(" +
		"id serial primary key, " +
		"tag varchar (50), " +
		"event_id integer, " +
		"email varchar (50), " +
		"name varchar (50), " +
		"password varchar (250), " +
		"created_at varchar (100), " +
		"last_access varchar (100), " +
		"solved_challenges text);"

	ADD_TEAM_QUERY = "INSERT INTO team (tag, event_id, email, name, password, created_at, last_access, solved_challenges)" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	ADD_EVENT_QUERY = "INSERT INTO event (tag, name, available, capacity, frontends, status, exercises, started_at, finish_expected)" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	UPDATE_EVENT_FINISH_DATE       = "UPDATE event SET finished_at = $2 WHERE tag = $1"
	UPDATE_EVENT_STATUS            = "UPDATE event SET status = $2 WHERE tag = $1 "
	UPDATE_EVENT_LASTACCESSED_DATE = "UPDATE team SET last_access = $2 WHERE tag = $1"
	UPDATE_TEAM_SOLVED_CHL         = "UPDATE team SET solved_challenges = $2 WHERE tag = $1"

	QUERY_SOLVED_CHLS = "SELECT solved_challenges FROM team WHERE tag=$1"
	QUERY_EVENT_TABLE = "SELECT * FROM event"

	QUERY_EVENT_ID    = "SELECT id FROM event WHERE tag=$1 and finished_at is null"
	QUERY_EVENT_TEAMS = "SELECT * FROM team WHERE event_id=$1"

	QUERY_EVENT_STATUS = "SELECT status FROM event WHERE tag=$1"
)
