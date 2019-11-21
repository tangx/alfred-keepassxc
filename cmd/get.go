package cmd

import (
	"encoding/json"
	"regexp"

	"github.com/tangx/alfred-keepassxc/keepassxc"
	"github.com/tobischo/gokeepasslib/v2"
)

// Get return json string
func (c Client) Get(reg string) (result string) {

	db := keepassxc.NewClient(c.DBpath, c.DBpass)
	ch := make(chan gokeepasslib.Entry, 5)

	patt := regexp.MustCompile(reg)

	root := db.Content.Root
	wg.Add(1)
	go walkRoot(root, ch)
	go fillItems(ch, patt)
	wg.Wait()

	body, _ := json.Marshal(items)
	return string(body)
}

func walkRoot(root *gokeepasslib.RootData, ch chan gokeepasslib.Entry) {
	defer close(ch)
	walkGroup(root.Groups[0], ch)
}

func walkGroup(group gokeepasslib.Group, ch chan gokeepasslib.Entry) {

	for _, entry := range group.Entries {
		ch <- entry
	}
	for _, g := range group.Groups {
		walkGroup(g, ch)
	}
}

func match(entry gokeepasslib.Entry, patt *regexp.Regexp) bool {
	if patt.Match([]byte(entry.GetTitle())) {
		return true
	}
	if patt.Match([]byte(entry.GetContent("UserName"))) {
		return true
	}
	return false
}

func fillItems(ch chan gokeepasslib.Entry, patt *regexp.Regexp) {
	defer wg.Done()
	for entry := range ch {
		if !match(entry, patt) {
			continue
		}

		ModShift := keepassxc.KeepassXCItem{
			Subtitle: entry.GetContent("UserName"),
			Arg:      entry.GetContent("UserName"),
			Valid:    false,
		}
		ModCmd := keepassxc.KeepassXCItem{
			Subtitle: entry.GetPassword(),
			Arg:      entry.GetPassword(),
			Valid:    false,
		}
		var item = keepassxc.KeepassXCItem{Valid: true,
			Title:    entry.GetTitle(),
			Arg:      entry.GetPassword(),
			Subtitle: entry.GetContent("UserName"),
			Mods: map[string]keepassxc.KeepassXCItem{
				"shift": ModShift,
				"cmd":   ModCmd,
			},
		}
		items = append(items, item)
	}
}
