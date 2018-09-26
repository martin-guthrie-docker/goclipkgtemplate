package goclipkgtemplate

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
)

// initializer struct for ExAClass
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

	// data from the config file
	if v.Get("Manufacturer") != nil {
		t.CarData.Manufacturer = v.Get("Manufacturer").(string)
	} else {
		t.CarData.Manufacturer = "UNKNOWN"
	}
	if v.Get("Model") != nil {
		t.CarData.Model = v.Get("Model").(string)
	} else {
		t.CarData.Model = "UNKNOWN"
	}
	//if v.Get("Year") != nil {
	//	t.CarData.Year = v.Get("Year").(int16)
	//} else {
	//	t.CarData.Year = -1
	//}
	//t.CarData.Options = v.Get("Options").(string)  # FIXME

	// data from the environment variables
	if viper.Get("one") != nil {
		t.EnvData.one = viper.Get("one").(string)
	} else {
		t.EnvData.one = "UNKNOWN"
	}

	t.Log.Info("Created")
	return t, nil
}

func (t *ConfigClass) Open() error {
	return nil
}

func (t *ConfigClass) Dump(saveToYaml bool, filename string) error {
	t.Log.Info("Manufacturer: ", t.CarData.Manufacturer)
	t.Log.Info("Model: ", t.CarData.Model)
	//t.Log.Info("Year: ", t.CarData.Year)

	t.Log.Info("ENV ONE: ", t.EnvData.one)

	if saveToYaml {
		t.Log.Info("TODO, save to file ", filename)
	}
	return nil
}
