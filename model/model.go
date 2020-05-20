package model

type Event struct {
	Tag                string
	Name               string
	Frontends          string
	Exercises          string
	Available          uint
	Capacity           uint
	StartedAt          string
	ExpectedFinishTime string
	FinishedAt         string
}

type Team struct {
	Id               string
	Email            string
	Name             string
	Password         string
	CreatedAt        string
	LastAccess       string
	SolvedChallenges string
}

type Config struct {
	Host 		string `yaml:"host,omitempty"`
	AuthKey 	string `yaml:"auth-key,omitempty"`
	SigninKey 	string `yaml:"signin-key,omitempty"`
	DB struct {
		Host string `yaml:"host,omitempty"`
		User string `yaml:"user,omitempty"`
		Pass string `yaml:"pass,omitempty"`
		Name string `yaml:"db_name,omitempty"`
		Port uint 	`yaml:"db_port,omitempty"`
	} `yaml:"db,omitempty"`
	TLS struct {
		Enabled bool `yaml:"enabled"`
		CertFile 	string `yaml:"certfile"`
		CertKey 	string `yaml:"certkey"`
		CAFile 		string `yaml:"cafile"`
	} `tls:"db,omitempty"`
}
