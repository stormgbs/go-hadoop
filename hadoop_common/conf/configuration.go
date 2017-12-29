package conf

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const (
	HADOOP_CONF_DIR = "HADOOP_CONF_DIR"
	// CORE_DEFAULT = "core-default.xml"
	CORE_SITE    = "core-site.xml"
	CORE_DEFAULT = CORE_SITE
	// HDFS_DEFAULT = "hdfs-default.xml"
	HDFS_SITE    = "hdfs-site.xml"
	HDFS_DEFAULT = HDFS_SITE
)

type Configuration interface {
	Get(key string, defaultValue string) (string, error)
	GetInt(key string, defaultValue int) (int, error)

	Set(key string, value string) error
	SetInt(key string, value int) error
}

type configuration struct {
	Properties map[string]string
}

type property struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type hadoopConfiguration struct {
	XMLName    xml.Name   `xml:"configuration"`
	Properties []property `xml:"property"`
}

func (conf *configuration) Get(key string, defaultValue string) (string, error) {
	value, exists := conf.Properties[key]
	if !exists {
		return defaultValue, nil
	}
	return value, nil
}

func (conf *configuration) GetInt(key string, defaultValue int) (int, error) {
	value, exists := conf.Properties[key]
	if !exists {
		return defaultValue, nil
	}
	return strconv.Atoi(value)
}

func (conf *configuration) Set(key string, value string) error {
	conf.Properties[key] = value
	return nil
}

func (conf *configuration) SetInt(key string, value int) error {
	conf.Properties[key] = strconv.Itoa(value)
	return nil
}

func NewConfiguration() (Configuration, error) {
	return NewConfigurationResources([]string{})
}

func NewConfigurationResources(resources []string) (Configuration, error) {
	// Add $HADOOP_CONF_DIR/core-default.xml & $HADOOP_CONF_DIR/core-site.xml
	resourcesWithDefault := []string{CORE_DEFAULT, CORE_SITE}
	resourcesWithDefault = append(resourcesWithDefault, resources...)

	c := configuration{Properties: make(map[string]string)}

	for _, resource := range resourcesWithDefault {
		conf, err := os.Open(os.Getenv(HADOOP_CONF_DIR) + string(os.PathSeparator) + resource)
		if err != nil {
			log.Fatal("Couldn't open resource: ", err)
			return nil, err
		}
		confData, err := ioutil.ReadAll(conf)
		if err != nil {
			log.Fatal("Couldn't read resource: ", err)
			return nil, err
		}
		defer conf.Close()

		// Parse
		var hConf hadoopConfiguration
		err = xml.Unmarshal(confData, &hConf)
		if err != nil {
			log.Fatal("Couldn't parse core-site.xml", err)
			return nil, err
		}

		// Save into configuration
		for _, kv := range hConf.Properties {
			c.Set(kv.Name, kv.Value)
		}
	}

	return &c, nil
}
