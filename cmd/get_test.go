package cmd

import (
	"fmt"
	"testing"
)

func Test_ALL(t *testing.T) {
	c := Client{
		DBpath: dbpath,
		DBpass: dbpass,
	}

	// fmt.Println("HELLO")
	result := c.Get("evernote")
	fmt.Println(result)
}
