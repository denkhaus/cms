package engine

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////
type FieldRecord struct {
	Id   bson.ObjectId `bson:"_id" json:"_id"`
	Name string        `bson:"nm" json:"nm"`
}

////////////////////////////////////////////////////////////////////////////////
func (e *Engine) IsFieldPresent(name string) (bool, error) {
	sess, coll := e.GetSystemSession("fields")
	defer sess.Close()

	query := coll.Find(bson.M{"nm": name})
	if cnt, err := query.Count(); err != nil {
		return false, fmt.Errorf("Types::IsFieldPresent::%s", err.Error())
	} else {
		return cnt > 0, nil
	}

	return false, nil
}

////////////////////////////////////////////////////////////////////////////////
func (e *Engine) FieldCreate(typeField, tp string) error {
	if ok, err := e.IsTypePresent(typeName); err != nil {
		return err
	} else if !ok {
		return fmt.Errorf("Types::FieldCreate::Type %s is not present", name)
	}

	if ok, err := e.IsFieldPresent(fieldName); err != nil {
		return err
	} else if !ok {
		return fmt.Errorf("Types::FieldCreate::Type %s is not present", name)
	}

	sess, coll := e.GetSystemSession("fields")
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
