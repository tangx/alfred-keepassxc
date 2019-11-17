package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/tangx/alfred-keepassxc/keepassxc"
	"github.com/tobischo/gokeepasslib/v2"
)

var kpa = keepassxc.KeepassXC{}

var pat = strings.Join(os.Args[1:], ".*")
var reg = regexp.MustCompile(pat)

func main() {
	// Login
	dbpath := os.Getenv("KPA_KDBX")
	dbpass := os.Getenv("KPA_PASS")
	db := keepassxc.NewClient(dbpath, dbpass)

	rootGroup := db.Content.Root.Groups[0]
	walkGroups(rootGroup)

	kpaByte, _ := json.Marshal(kpa)
	fmt.Printf("%s", kpaByte)
}

func walkGroups(group gokeepasslib.Group) {
	for _, entry := range group.Entries {
		addItems(entry)
	}
	for _, subGroup := range group.Groups {
		walkGroups(subGroup)
	}
}

func addItems(entry gokeepasslib.Entry) {

	title := entry.GetTitle()
	userName := entry.GetContent("UserName")

	if reg.Match([]byte(title)) || reg.Match([]byte(userName)) {
		kpaItem := keepassxc.KeepassXCItem{
			Valid:    true,
			Title:    title,
			Arg:      entry.GetPassword(),
			Subtitle: entry.GetContent("UserName"),
		}

		kpa.Items = append(kpa.Items, kpaItem)
	}
}
