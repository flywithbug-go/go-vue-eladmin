package log_writer

import (
	"sync"

	"github.com/flywithbug/log4go"
)

const tunnelSizeDefault = 1024

var (
	DBlogPool       *sync.Pool
	dbWriterDefault *DBWriter
	takeup          = false
)

func init() {
	dbWriterDefault = NewDBWriter()
	DBlogPool = &sync.Pool{New: func() interface{} {
		return new(Log)
	}}
}

type DBWriter struct {
	tunnel chan *Log
	c      chan bool
}

func NewDBWriter() *DBWriter {
	if dbWriterDefault != nil && takeup == false {
		takeup = true
		return dbWriterDefault
	}
	dw := new(DBWriter)
	dw.tunnel = make(chan *Log, tunnelSizeDefault)
	dw.c = make(chan bool, 1)
	go bootstrapLogWriter(dw)
	return dw
}

func (db *DBWriter) Init() error {
	return nil
}

func (db *DBWriter) Write(record *log4go.Record) error {
	l, ok := record.Ext.(*Log)
	if !ok {
		l = DBlogPool.Get().(*Log)
	}
	l.Info = record.Info
	l.Code = record.Code
	l.Time = record.Time
	l.Level = record.Level
	l.Flag = LEVEL_FLAGS[record.Level]
	db.tunnel <- l
	return nil
}

func bootstrapLogWriter(db *DBWriter) {
	if db == nil {
		panic("writer is nil")
	}
	var (
		ok bool
	)
	if _, ok = <-db.tunnel; !ok {
		db.c <- true
		return
	}

	for {
		select {
		case l, ok := <-db.tunnel:
			if !ok {
				db.c <- true
				return
			}
			l.Insert()
			DBlogPool.Put(l)
		}
	}
}
