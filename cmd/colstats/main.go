package main

import (
	"fmt"
	"github.com/mchirico/go_network/ssh"
	"github.com/mchirico/go_network/yamlpkg"
	"log"
	"os"
)

func main() {
	home := os.Getenv("HOME")

	file := home + "/.networkScriptConfig.yaml"
	fmt.Println(file)

	c := yamlpkg.Config{}
	err := c.Read(file)
	if err != nil {
		log.Fatalf("We're not on a system that can test ssh.: %v\n", err)
		return
	}

	idx := 3

	s := ssh.SSH{}
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

	for idx, v := range c.Yaml[idx].Config.IP {
		s.Server = v
		fmt.Println(idx, v)
		s.CmdServers()
	}
}
