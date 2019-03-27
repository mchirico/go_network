package ssh

import (
	"fmt"
	"github.com/mchirico/go_network/yamlpkg"
	"os"
	"testing"
)

func TestSSH_Example(t *testing.T) {

	home := os.Getenv("HOME")

	file := home + "/.networkScriptConfig.yaml"
	fmt.Println(file)

	c := yamlpkg.Config{}
	err := c.Read(file)
	if err != nil {
		t.Logf("We're not on a system that can test ssh.")
		return
	}

	idx := 2

	s := SSH{}
	s.Password = c.Yaml[idx].Config.Password
	s.Server = c.Yaml[idx].Config.IP[0]
	s.User = c.Yaml[idx].Config.Username
	s.Password = c.Yaml[idx].Config.Password
	s.File = c.Yaml[idx].Config.FileOut
	s.CMD = c.Yaml[idx].Config.Cmd
	s.UseSSHkey = c.Yaml[idx].Config.UseSSHkey
	s.SSHPubKey = c.Yaml[idx].Config.SSHPubKey
	if c.Yaml[idx].Config.Repeats == 0 {
		s.Repeats = 1
	} else {
		s.Repeats = c.Yaml[idx].Config.Repeats
	}

	fmt.Println(c.Yaml[idx].Config.Repeats)

	s.CmdServers()

}
