package cmd

import (
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

var items []keepassxc.KeepassXCItem
var wg sync.WaitGroup
