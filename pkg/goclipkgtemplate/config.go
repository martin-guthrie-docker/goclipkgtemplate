package goclipkgtemplate

import (
	"fmt"
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

	err := v.Unmarshal(&t.CarData)
	if err != nil {
		panic(err)
	}

	// data from the environment variables
	// was not able to use Unmarshal.... see https://github.com/spf13/viper/issues/188
	if v.Get("one") != nil {
		t.EnvData.one = v.Get("one").(string)
	} else {
		t.EnvData.one = "UNKNOWN"
	}

	return t, nil
}

func (t *ConfigClass) Dump(toConsole bool, saveToYaml bool, filename string) error {
	t.Log.Infof("CarData: %s %s %d %s",
		t.CarData.Manufacturer, t.CarData.Model, t.CarData.Year, t.CarData.Options)
	if toConsole {
		fmt.Printf("CarData: %s %s %d %s\n",
			t.CarData.Manufacturer, t.CarData.Model, t.CarData.Year, t.CarData.Options)
	}

	t.Log.Infof("ENV GOCLIP_ONE: %s", t.EnvData.one)
	if toConsole {
		fmt.Printf("ENV GOCLIP_ONE: %s\n", t.EnvData.one)
	}

	if saveToYaml {
		t.Log.Info("TODO, save to file ", filename)
	}
	return nil
}
