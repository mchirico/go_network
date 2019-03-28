package yamlpkg

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func WriteYaml(file string) {
	c := Config{}

	ips := []string{"0.0.0.0", "0.0.0.1"}

	y := y{ips, "uptime;hostname;iptables -nvL",
		"Tom",
		"password123",
		true,
		"id_ed25519",
		"file0",
		2}
	g := group{"Group0", y}
	g1 := group{"Group1", y}

	gg := []group{g, g1}

	c.Yaml = gg

	//c.Yaml[0].Config.Password
	fmt.Println(c.Yaml[1].Config.Repeats)

	c.Write(file)

}

func TestReadWrite(t *testing.T) {
	file := ".networkScriptConfig.yaml"
	WriteYaml(file)

	c := Config{}
	c.Read(file)

	if c.Yaml[1].Group != "Group1" {
		t.Fatalf("Not able to read group")
	}

}

func TestConfig_ListGroups(t *testing.T) {
	file := ".networkScriptConfig.yaml"
	WriteYaml(file)

	c := Config{}
	c.Read(file)

	if c.ListGroups()[0] != "Group0" {
		t.Fatalf("Expected: Group0, got: %s\n",
			c.ListGroups()[0])
	}
	if c.ListGroups()[1] != "Group1" {
		t.Fatalf("Expected: Group1, got: %s\n",
			c.ListGroups()[1])
	}
}
