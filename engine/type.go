package engine

import (
	"fmt"

	"github.com/denkhaus/tcgl/applog"
	"gopkg.in/mgo.v2/bson"
)

type Type struct {
	TypeRecord
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
type TypeRecord struct {
	Id   bson.ObjectId `bson:"_id" json:"_id"`
	Name string        `bson:"nm" json:"nm"`
}

////////////////////////////////////////////////////////////////////////////////
func (e *Engine) IsTypePresent(name string) (bool, error) {
	sess, coll := e.GetSystemSession("types")
	defer sess.Close()

	query := coll.Find(bson.M{"nm": name})
	if cnt, err := query.Count(); err != nil {
		return false, fmt.Errorf("Types::IsTypePresent::%s", err.Error())
	} else {
		return cnt > 0, nil
	}

	return false, nil
}

////////////////////////////////////////////////////////////////////////////////
func (e *Engine) TypeCreate(name string) error {
	if ok, err := e.IsTypePresent(name); err != nil {
		return err
	} else if ok {
		return fmt.Errorf("Types::TypeCreate::Type %s is already present", name)
	}

	sess, coll := e.GetSystemSession("types")
	defer sess.Close()

	rec := TypeRecord{
		Id:   bson.NewObjectId(),
		Name: name,
	}

	if err := coll.Insert(&rec); err != nil {
		return fmt.Errorf("Types::TypCreate::Insert::%s", err.Error())
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
func (e *Engine) TypeRemove(name string) error {
	if ok, err := e.IsTypePresent(name); err != nil {
		return err
	} else if !ok {
		return fmt.Errorf("Types::TypeRemove::Type %s is not present", name)
	}

	sess, coll := e.GetSystemSession("types")
	defer sess.Close()

	if _, err := coll.RemoveAll(bson.M{"nm": name}); err != nil {
		return fmt.Errorf("Types::TypeRemove::Remove::%s", err.Error())
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
func (e *Engine) TypeList() error {
	sess, coll := e.GetSystemSession("types")
	defer sess.Close()

	iter := coll.Find(nil).Iter()
	if err := iter.Err(); err != nil {
		return fmt.Errorf("Types::Iter::Error::%s", err.Error())
	}

	rec := TypeRecord{}
	for iter.Next(&rec) {
		applog.Infof("%s", rec.Name)
	}

	if err := iter.Close(); err != nil {
		return fmt.Errorf("Types::Iter::Close::%s", err.Error())
	}

	return nil
}
