package main

import (
	"fmt"
	"github.com/koomox/ext"
	"github.com/koomox/static"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func parseArgs() (root, prefix, filter, output string, err error) {
	for _, arg := range os.Args {
		op := strings.Split(arg, "=")
		sc := op[0]
		switch len(op) {
		case 2:
			switch sc {
			case "--root":
				root = strings.Replace(op[1], "\\", "/", -1)
			case "--prefix":
				prefix = strings.Replace(op[1], "\\", "/", -1)
			case "--filter":
				filter = strings.Replace(op[1], "\\", "/", -1)
			case "--output", "--out":
				output = op[1]
			}
		}
	}
	switch {
	case root == "":
		if root , err= ext.GetCurrentDirectory(); err != nil {
			return
		}
	case output == "":
		output = path.Join(root, "files.compress.txt")
	}
	return
}

func main() {
	root, prefix, filter, output, err := parseArgs()
	if err != nil {
		fmt.Println(err)
		return
	}
	enc := static.NewEncoding()
	enc.New().Root(root).Prefix(prefix).Filter(filter)
	buf, err := enc.CompressAllFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = ioutil.WriteFile(output, buf, os.ModePerm); err != nil {
		fmt.Println(err)
	}
}