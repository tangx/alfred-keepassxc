package keepassxc

import (
	"os"

	"github.com/tangx/alfred-keepassxc/utils"
	"github.com/tobischo/gokeepasslib/v2"
)

type KeepassXC struct {
	Items []KeepassXCItem        `json:"items,omitempty"`
	DB    *gokeepasslib.Database `json:"db,omitempty"`
}

type KeepassXCItem struct {
	Arg      string                   `json:"arg,omitempty"`
	Mods     map[string]KeepassXCItem `json:"mods,omitempty"`
	Subtitle string                   `json:"subtitle,omitempty"`
	Title    string                   `json:"title,omitempty"`
	Valid    bool                     `json:"valid,omitempty"`
}

// NewClient return a new keepass clien
func NewClient(dbpath string, dbpass string) *gokeepasslib.Database {

	file, err := os.Open(dbpath)
	utils.IsError(err)

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials(dbpass)
	_ = gokeepasslib.NewDecoder(file).Decode(db)

	err = db.UnlockProtectedEntries()
	utils.IsError(err)
	return db
}
