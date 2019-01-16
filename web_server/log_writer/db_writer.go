package log_writer

import "github.com/flywithbug/log4go"

type DBWriter struct {
	DBName     string `json:"db_name"`
	Collection string `json:"collection"`
}

func NewDBWriter() *DBWriter {
	return &DBWriter{}
}

func (db *DBWriter) Init() error {
	return nil
}

func (db *DBWriter) Write(*log4go.Record) error {

	return nil
}
