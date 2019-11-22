package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tangx/alfred-keepassxc/keepassxc"
	"github.com/tangx/alfred-keepassxc/utils"
)

const (
	lower = `abcdefghijkllmnopqrstuvwxyz`
	num   = `0123456789`
	alpha = `!@%^*()_+.`
)

// Password return a password
func Password(n int, complex bool) string {
	rand.Seed(time.Now().UnixNano())
	upper := strings.ToUpper(lower)

	pool := fmt.Sprintf("%s%s%s", lower, upper, num)
	if complex {
		pool += alpha
	}

	password := make([]byte, n)
	for i := 0; i < n; i++ {
		r := rand.Intn(len(pool))
		password[i] = pool[r]
	}

	return string(password)
}

// GenUUID return a uuid
func GenUUID() {

	UUID := uuid.New().String()
	var item = keepassxc.KeepassXCItem{
		Valid:    true,
		Title:    UUID,
		Arg:      UUID,
		Subtitle: "复制",
	}
	items = append(items, item)
}

// GenPassword return a []keepassxc.KeepassXCItem
func GenPassword(n int, complex bool) {
	comp := "简单"
	if complex {
		comp = "复杂"
	}
	subtitle := fmt.Sprintf("复制 %d位 %s密码", n, comp)
	password := Password(n, complex)
	var item = keepassxc.KeepassXCItem{
		Valid:    true,
		Title:    password,
		Arg:      password,
		Subtitle: subtitle,
	}
	items = append(items, item)
}

// Gen a password
func Gen() {
	GenPassword(12, false)
	GenPassword(12, true)
	GenPassword(16, false)
	GenPassword(16, true)
	GenUUID()

	PrintResult()
}

// PrintResult print out json result string
func PrintResult() {
	body, err := json.Marshal(items)
	utils.IsError(err)
	fmt.Println(string(body))
}
