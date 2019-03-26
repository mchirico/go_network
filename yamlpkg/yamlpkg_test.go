package yamlpkg

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestReadWrite(t *testing.T) {
	file := ".networkScriptConfig.yaml"

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

	c.Write(file)

	c2 := Config{}
	c2.Read(file)

}
