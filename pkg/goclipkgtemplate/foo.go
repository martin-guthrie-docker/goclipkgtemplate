package goclipkgtemplate

import (
	"fmt"
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
	msg := fmt.Sprintf( "Hello! My name is %s", t.Name)
	t.Log.Info(msg)
	fmt.Println(msg)
	return nil
}

func (t *FooClass) IsOpen() bool {
	return t.open
}

func (t *FooClass) Close() {
	msg := fmt.Sprintf( "And my host is %s", t.Host)
	t.Log.Info(msg)
	fmt.Println(msg)
	t.open = false
}