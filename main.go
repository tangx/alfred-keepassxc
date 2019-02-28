package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/tobischo/gokeepasslib/v2"
)

// Keepassxc result Json Struct
type Keepassxc struct {
	Items []KeepassxcItem `json:"items"`
}

// KeepassxcItem Keepassxc's Items Json Struct
type KeepassxcItem struct {
	Arg      string `json:"arg"`
	Subtitle string `json:"subtitle"`
	Title    string `json:"title"`
	Valid    bool   `json:"valid"`
}

var kpa = Keepassxc{}
var reg = regexp.MustCompile(os.Args[1])

func main() {

	// Login
	dbfile := os.Getenv("KPA_KDBX")
	// keyfile := "/Users/tangxin/kp.key"
	dbpass := os.Getenv("KPA_PASS")

	// fmt.Println(dbfile)
	// fmt.Println(dbpass)

	file, _ := os.Open(dbfile)
	defer file.Close()

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials(dbpass)

	_ = gokeepasslib.NewDecoder(file).Decode(db)
	db.UnlockProtectedEntries()

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

	// ok, _ := regexp.MatchString("www", title)

	reg.Match([]byte(title))

	if reg.Match([]byte(title)) {

		kpaItem := KeepassxcItem{
			Valid:    true,
			Title:    title,
			Arg:      entry.GetPassword(),
			Subtitle: entry.GetContent("UserName"),
		}
		kpa.Items = append(kpa.Items, kpaItem)
	}
}
