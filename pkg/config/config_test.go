package config

import (
	"fmt"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	r, _ := os.Open("config.yaml")

	fmt.Println(GetConfigHome())

	c, err := Parse(r)
	if err != nil {
		panic(err.Error())
	}
	c.AddContext("admin@local", "remote", "admin")
	c.AddUser("admin", "admin", "admin")
	c.AddCluster("remote", "remote", 5672)

	c.SwitchContext("admin@local")

	c.Print()
	fmt.Println(c.GetCurrentContextDsn())

	fmt.Printf("%+v", c)
}
