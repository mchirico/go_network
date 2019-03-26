package yamlpkg

// REF: http://sweetohm.net/article/go-yaml-parsers.en.html

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

type y struct {
	IP       []string `yaml:"ips"`
	Cmd      string   `yaml:"cmd"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	UseSSHkey   bool   `yaml:"usesshkey"`
	SSHPubKey string `yaml:"sshpubkeyfile"`
	FileOut  string   `yaml:"fileOut"`
	Repeats  int       `yaml:"commandrepeats"`
}

type group struct {
	Group  string `yaml:"group"`
	Config y      `yaml:"config"`
}

// Config entry point
type Config struct {
	sync.Mutex
	Yaml []group
}

// Write yaml file
func (c *Config) Write(file string) error {
	c.Lock()
	defer c.Unlock()

	data, err := yaml.Marshal(c.Yaml)
	if err != nil {
		log.Printf("yaml.Marshal(config): %v", err)
		return err
	}

	err = ioutil.WriteFile(file, data, 0600)
	if err != nil {
		log.Printf("error in yaml write: %v", err)
	}
	return err
}

// Read yaml file
func (c *Config) Read(file string) error {
	c.Lock()
	defer c.Unlock()

	source, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("Error ioutil.ReadFile")
		return err
	}
	err = yaml.Unmarshal(source, &c.Yaml)
	if err != nil {
		log.Printf("Error Unmarshal")
		return err
	}

	return err
}

// SetDefault simple config settings
func (c *Config) SetDefault() {
	c.Lock()
	defer c.Unlock()

	ips := []string{"0.0.0.0", "0.0.0.1"}

	y := y{ips, "date;hostname;uptime",
		"Tom",
		"p@ssworD",
		true,
		"id_ed25519",
		"fileOut.txt",
	2}
	g := group{"Group0", y}
	g1 := group{"Group1", y}

	gg := []group{g, g1}

	c.Yaml = gg

}
