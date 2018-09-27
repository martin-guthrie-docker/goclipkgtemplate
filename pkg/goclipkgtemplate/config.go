package goclipkgtemplate

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
)

// initializer struct for ConfigClass
type ConfigClassCfg struct {
	Log                  *logrus.Logger
}

type carData struct {
	Manufacturer         string
	Model                string
	Year                 int16
	Options              []string
}

type envData struct {
	one                  string
}

type ConfigClass struct {
	ConfigClassCfg       // this is an embedded type
	CarData              carData
	EnvData              envData
}

// constructor for ConfigClass
func NewConfigClass(v *viper.Viper, cfg ConfigClassCfg) (*ConfigClass, error) {

	// if no logger, create a null logger
	if cfg.Log == nil {
		cfg.Log = logrus.New()
		cfg.Log.Out = ioutil.Discard
	}

	t := new(ConfigClass)
	t.Log = cfg.Log

	err := viper.Unmarshal(&t.CarData)
	if err != nil {
		panic(err)
	}

	// data from the environment variables
	if viper.Get("one") != nil {
		t.EnvData.one = viper.Get("one").(string)
	} else {
		t.EnvData.one = "UNKNOWN"
	}

	t.Log.Info("Created")
	return t, nil
}

func (t *ConfigClass) Dump(saveToYaml bool, filename string) error {
	t.Log.Infof("CarData: %s %s %d %s",
		t.CarData.Manufacturer, t.CarData.Model, t.CarData.Year, t.CarData.Options)

	t.Log.Info("ENV ONE: ", t.EnvData.one)

	if saveToYaml {
		t.Log.Info("TODO, save to file ", filename)
	}
	return nil
}
