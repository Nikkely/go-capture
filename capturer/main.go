package capturer

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// type Config struct {
// 	CapConf []CapConf
// }

func Run(path string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var config CapConf
	if err := yaml.Unmarshal(buf, &config); err != nil {
		return err
	}

	return Capture(config)
}
