package tool

import (
	"fengqi/qbittorrent-tool/config"
	"fengqi/qbittorrent-tool/qbittorrent"
	"fmt"
	"strings"
)

func StatusTag(c *config.Config, torrent *qbittorrent.Torrent) {
	if !c.StatusTag.Enable || c.StatusTag.MapConfig == nil {
		return
	}

	trackerList, err := qbittorrent.Api.GetTorrentTrackers(torrent.Hash)
	if err != nil || len(trackerList) == 0 {
		fmt.Printf("[ERR] get %s tracker list err: %v, count: %d\n", torrent.Name, err, len(trackerList))
		return
	}

	tag := ""
	miss := make(map[string]int, 0)
	for _, tracker := range trackerList {
		if tracker.Status == 2 || tracker.Msg == "" {
			return
		}

		matched := false
		for key, value := range c.StatusTag.MapConfig {
			if strings.Contains(tracker.Msg, key) {
				tag = value
				matched = true
				break
			}
		}

		if !matched {
			miss[tracker.Msg] += 1
		}
	}

	if len(miss) > 0 {
		for item, _ := range miss {
			fmt.Printf("err: \"%s: %s\" not map config\n", torrent.Name, item)
		}
	}

	if tag == "" || strings.Contains(torrent.Tags, tag) {
		return
	}

	err = qbittorrent.Api.AddTags(torrent.Hash, tag)
	if err != nil {
		fmt.Printf("[ERR] add tag %s to %s err: %v\n", tag, torrent.Name, err)
		return
	}

	fmt.Printf("[INFO] add tag %s to %s\n", tag, torrent.Name)
}
