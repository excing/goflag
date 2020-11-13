package goflag

import (
	"fmt"
	"os"
	"testing"
)

type Config struct {
	Debug bool
}

type Config01 struct {
	Debug bool `flag:"Is debug"`
}

func init() {
	fmt.Fprintln(os.Stdout, "hello goflag")
}

func TestFlagErr(t *testing.T) {
	var debug = false
	Var(&debug, "debug")
	Parse("config", "Test config")
}

func TestFlagSuc(t *testing.T) {
	var debug = false
	Var(&debug, "debug", "Is debug")
	Parse("config", "Test config")
}

func TestConfigErr(t *testing.T) {
	var config Config
	Var(&config)
	Parse("config", "Test Config")
}

func TestConfigSuc(t *testing.T) {
	var config Config01
	Var(&config)
	Parse("config", "Test Config")
}
