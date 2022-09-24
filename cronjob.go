package main

import (
	"errors"
	"strings"
)

// represents a single cron entry i.e. a single line in a crontab
type cronjob struct {
	timespec string
	command  string
}

// create a new 'cronjob' struct based on the given line from a crontab
func NewCronjob(jobspec string) (*cronjob, error) {
	words := strings.Fields(jobspec)

	if len(words) < 1 || len(words[0]) == 0 {
		return nil, errors.New("Invalid jobspec: " + jobspec)
	}

	if words[0][0:1] == "@" {
		if len(words) < 2 {
			return nil, errors.New("Invalid jobspec: " + jobspec)
		}
		return &cronjob{
			timespec: strings.Join(words[0:1], " "),
			command:  strings.Join(words[1:], " "),
		}, nil
	} else {
		if len(words) < 6 {
			return nil, errors.New("Invalid jobspec: " + jobspec)
		}
		return &cronjob{
			timespec: strings.Join(words[0:5], " "),
			command:  strings.Join(words[5:], " "),
		}, nil
	}
}
