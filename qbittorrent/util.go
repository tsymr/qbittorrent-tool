package qbittorrent

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func (t *Torrent) GetTrackerHost() (string, error) {
	if t.Tracker == "" {
		return "", errors.New("tracker empty")
	}

	parse, err := url.Parse(t.Tracker)
	if err != nil {
		return "", err
	}

	split := strings.Split(parse.Host, ".")
	if len(split) < 2 {
		return parse.Host, nil
	}
	return fmt.Sprintf("%s.%s", split[len(split)-2], split[len(split)-1]), nil
}
