package mongo

import (
	"errors"
	"gopkg.in/mgo.v2"
)

var sessionMap = make(map[string]*sessionMgo)


type Mongoer interface {
	IsDestroyed() bool
	Destroy()
	Use(alias string) error
	UseDB(db string)
	Session()*mgo.Session
	CurrentDB() string
	Collection(c string) *mgo.Collection
	// CURD
	Insert(collection string, docs ...interface{}) error
	Update(collection string, selector interface{}, update interface{}) error
	Find(collection string, query interface{}, results interface{}) error
	FindOne(collection string, query interface{}, result interface{}) error
	Remove(collection string, selector interface{}) error
}

type sessionMgo struct {
	db 			string
	destroyed  	bool
	session  	*mgo.Session
}

func (s *sessionMgo)IsDestroyed()bool  {
	return s.destroyed
}
func (s *sessionMgo)Destroy()  {
	if s.destroyed {
		return
	}
	s.destroyed = true
	s.session.Close()
}


func (s *sessionMgo) Use(alias string) error {
	if s.destroyed {
		return errors.New(ErrMongoObjDestroyed)
	}
	if alias == ""{
		alias = "default"
	}
	c, ok := sessionMap[alias]
	if ok {
		s.db = c.db
		s.session  = c.session.Clone()
		return nil
	}
	return errors.New(ErrNoConnection + " named " + alias)
}

func (s *sessionMgo) UseDB(db string) {
	if db == "" {
		return
	}
	s.db = db
}

func (s *sessionMgo) Session() *mgo.Session {
	if s.destroyed {
		return nil
	}
	return s.session
}

func (s *sessionMgo) CurrentDB() string {
	return s.db
}

func (s *sessionMgo) Collection(c string) *mgo.Collection {
	if s.destroyed {
		return nil
	}
	if c == "" {
		return nil
	}
	return s.session.DB(s.db).C(c)
}

func (s *sessionMgo) Insert(collection string, docs ...interface{}) error {
	if s.destroyed {
		return errors.New(ErrMongoObjDestroyed)
	}
	c := s.Collection(collection)
	if c == nil {
		return errors.New(ErrCannotSwitchCollection + " '" + collection + "' in db '" + s.db + "'")
	}
	return c.Insert(docs)
}

func (s *sessionMgo) Update(collection string, selector interface{}, update interface{}) error {
	if s.destroyed {
		return errors.New(ErrMongoObjDestroyed)
	}
	c := s.Collection(collection)
	if c == nil {
		return errors.New(ErrCannotSwitchCollection + " '" + collection + "' in db '" + s.db + "'")
	}
	return c.Update(selector,update)
}

func (s *sessionMgo) Find(collection string, query interface{}, results interface{}) error {
	if s.destroyed {
		return errors.New(ErrMongoObjDestroyed)
	}
	c := s.Collection(collection)
	if c == nil {
		return errors.New(ErrCannotSwitchCollection + " '" + collection + "' in db '" + s.db + "'")
	}
	return c.Find(query).All(results)
}

func (s *sessionMgo) FindOne(collection string, query interface{}, result interface{}) error {
	if s.destroyed {
		return errors.New(ErrMongoObjDestroyed)
	}
	c := s.Collection(collection)
	if c == nil {
		return errors.New(ErrCannotSwitchCollection + " '" + collection + "' in db '" + s.db + "'")
	}
	return c.Find(query).One(result)
}

func (s *sessionMgo) Remove(collection string, selector interface{}) error {
	if s.destroyed {
		return errors.New(ErrMongoObjDestroyed)
	}
	c := s.Collection(collection)
	if c == nil {
		return errors.New(ErrCannotSwitchCollection + " '" + collection + "' in db '" + s.db + "'")
	}
	return c.Remove(selector)
}

