package cmd

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/tangx/alfred-keepassxc/keepassxc"
	"github.com/tobischo/gokeepasslib/v2"
)

// Client for KXC
type Client struct {
	DBpath string
	DBpass string
	DB     *gokeepasslib.Database `json:"db,omitempty"`
}

var kpcItems []keepassxc.KeepassXCItem
var wg sync.WaitGroup

// Main to starts
func Main(args []string) {
	switch args[0] {
	case "gen":
		{
			Gen()
		}
	case "get":
		{
			Get(args[1:])
		}
	}
}

func Printout() {
	items := map[string][]keepassxc.KeepassXCItem{
		"items": kpcItems,
	}
	body, _ := json.Marshal(items)
	fmt.Printf("%s", body)
}
