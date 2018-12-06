package mongo

import "gopkg.in/mgo.v2"

type sessionMgo struct {
	db 			string
	destroyed  	bool
	session  *mgo.Session
}

type Mongoer interface {
	IsDestroyed() bool
	Destroy()

}


func (s sessionMgo)IsDestroyed()bool  {
	return s.destroyed
}
func (s sessionMgo)Destroy()  {
	if s.destroyed {
		return
	}
	s.destroyed = true
	s.session.Close()
}
