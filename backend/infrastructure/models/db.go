package models

type PostgreSQL struct {
	Connstr string
	DBname  string
}

var PostgreSetting = &PostgreSQL{}
