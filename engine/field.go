package engine

import (
	"fmt"
	"regexp"

	"gopkg.in/mgo.v2/bson"
)

var TypeFieldRegex = regexp.MustCompile("^([a-zA-Z0-9]+):([a-zA-Z0-9]+)$")

var types = map[string]struct{}{
	"string": struct{}{},
}

type Field struct {
	FieldRecord
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
type FieldRecord struct {
	Id     bson.ObjectId `bson:"_id" json:"_id"`
	TypeId bson.ObjectId `bson:"tpid" json:"tpid"`
	Name   string        `bson:"nm" json:"nm"`
}

////////////////////////////////////////////////////////////////////////////////
func (t *Type) IsFieldPresent(name string) (bool, error) {
	sess, coll := t.Engine.GetSystemSession("fields")
	defer sess.Close()

	query := coll.Find(bson.M{"tpid": t.Id, "nm": name})
	if cnt, err := query.Count(); err != nil {
		return false, fmt.Errorf("Types::IsFieldPresent::%s", err.Error())
	} else {
		return cnt > 0, nil
	}

	return false, nil
}

////////////////////////////////////////////////////////////////////////////////
func (e *Engine) FieldCreate(fp, tp string) error {
	if err := matchFieldPattern(fp); err != nil {
		return err
	}

	if err := matchFieldType(tp); err != nil {
		return err
	}

	res := TypeFieldRegex.FindStringSubmatch(fp)
	if len(res) != 3 {
		return fmt.Errorf("Fields::FieldCreate::Unable to parse Type:Field info.")
	}

	t, err := e.TypeLoad(res[1])
	if err != nil {
		return err
	}

	if ok, err := t.IsFieldPresent(res[2]); err != nil {
		return err
	} else if ok {
		return fmt.Errorf("Fields::FieldCreate::Field '%s' is already present.", res[2])
	}

	sess, coll := e.GetSystemSession("fields")
	defer sess.Close()

	rec := FieldRecord{
		Id:     bson.NewObjectId(),
		TypeId: t.Id,
		Name:   res[2],
	}

	if err := coll.Insert(&rec); err != nil {
		return fmt.Errorf("Fields::FieldCreate::Insert::%s", err.Error())
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
func (e *Engine) FieldRemove(fp string) error {
	if err := matchFieldPattern(fp); err != nil {
		return err
	}

	res := TypeFieldRegex.FindStringSubmatch(fp)
	if len(res) != 3 {
		return fmt.Errorf("Fields::FieldList::Unable to parse Type:Field info.")
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
func (e *Engine) FieldList(fp string) error {
	if err := matchFieldPattern(fp); err != nil {
		return err
	}

	res := TypeFieldRegex.FindStringSubmatch(fp)
	if len(res) != 3 {
		return fmt.Errorf("Fields::FieldList::Unable to parse Type:Field info.")
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
func matchFieldType(tf string) error {
	if tf == "" {
		return fmt.Errorf("Fields::FieldType is not specified. Value is mandatory.")
	}

	if _, ok := types[tf]; !ok {
		return fmt.Errorf("Fields::FieldType %s is unknown.")
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
func matchFieldPattern(fp string) error {
	if !TypeFieldRegex.Match([]byte(fp)) {
		return fmt.Errorf("Fields::Type:Field combination 'tp' is not valid. Use e.g. 'picture:url'.")
	}

	return nil
}
