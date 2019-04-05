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
	WriteYamlTest(file)

	c := Config{}
	c.Read(file)

	if c.Yaml[1].Group != "Group1" {
		t.Fatalf("Not able to read group")
	}

}

func TestConfig_ListGroups(t *testing.T) {
	file := ".networkScriptConfig.yaml"
	WriteYamlTest(file)

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
