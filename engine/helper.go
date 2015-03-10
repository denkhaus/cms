package engine

import (
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/mgo.v2"
)

func Inspect(args ...interface{}) {
	spew.Dump(args)
}

type MongoSessionProvider struct {
	Session *mgo.Session
}

func (p MongoSessionProvider) GetSystemSession(collName string) (*mgo.Session, *mgo.Collection) {
	sess := p.Session.Copy()
	coll := sess.DB("cms").C(collName)
	return sess, coll
}
