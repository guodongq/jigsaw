package main

import (
	"github.com/guodongq/jigsaw/pkg/designer"
	"github.com/guodongq/jigsaw/pkg/util/profile"
)

func main() {
	defer profile.Profile().Stop()

	designer.Execute()
}
