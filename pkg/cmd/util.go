package cmd

import (
	"errors"
	"fmt"
	"strings"
)

func getFullDSN(user, password, dsn string) (string, error) {
	if password == "" && user == "" {
		return dsn, nil
	}
	if password == "" || user == "" {
		return "", errors.New("Need both user and password")
	}
	t := strings.Split(dsn, "@")
	return fmt.Sprintf("%s:%s@%s", user, password, t[len(t)-1]), nil
}
