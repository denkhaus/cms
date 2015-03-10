package engine

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

////////////////////////////////////////////////////////////////////////////////
type Engine struct {
	MongoSessionProvider
	config *Config
}
type EngineFunc func(engine *Engine) error

////////////////////////////////////////////////////////////////////////////////
func (e *Engine) Execute(fn EngineFunc) error {
	return fn(e)
}

////////////////////////////////////////////////////////////////////////////////
func NewEngine(cnf *Config) (*Engine, error) {
	e := &Engine{config: cnf}

	if mConf, err := cnf.GetMongoConfig(); err != nil {
		return nil, err
	} else {

		if session, err := mgo.Dial(mConf["host"].(string)); err != nil {
			return nil, fmt.Errorf("Init error:: Mongo Session could not be initialized. :: %s", err.Error())
		} else {
			session.SetMode(mgo.Monotonic, true)
			e.Session = session
		}
	}

	return e, nil
}
