package goclipkgtemplate

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

// initializer struct for ExAClass
type FooClassCfg struct {
	Name                 string
	Log                  *logrus.Logger
}

type FooClass struct {
	FooClassCfg          // this is an embedded type
	Host                 string
	open                 bool
}

// constructor for ExAClass
func NewFooClass(host string, cfg FooClassCfg) (*FooClass, error) {

	// if no logger, create a null logger
	if cfg.Log == nil {
		cfg.Log = logrus.New()
		cfg.Log.Out = ioutil.Discard
	}

	if len(cfg.Name) == 0 {
		cfg.Name = "Unknown"
	}

	t := new(FooClass)
	t.Name = cfg.Name
	t.Log = cfg.Log
	t.Host = host
	t.open = false

	return t, nil
}

func (t *FooClass) Open() error {
	t.open = true
	t.Log.Info("Hello! My name is ", t.Name)
	return nil
}

func (t *FooClass) Close() {
	t.Log.Info("Goodbye ", t.Host)
	t.open = false
}