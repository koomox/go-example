package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"github.com/koomox/ext"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type File struct {
	Name string
	Content []byte
}

func Compress(data []byte) (buf []byte, err error){
	var b bytes.Buffer
	w := zlib.NewWriter(&b)

	if _, err = w.Write(data);err != nil {
		return
	}

	if err = w.Close(); err != nil {
		return
	}

	buf = b.Bytes()
	return
}

func TarCompressAllFile(root, prefix string) (buffer []byte, err error) {
	var (
		fs []string
		b bytes.Buffer
		buf []byte
	)
	if root == "" {
		if root, err = ext.GetCurrentDirectory(); err != nil {
			return
		}
	}

	root = path.Join(root, "")
	if fs, err = ext.GetCustomDirectoryAllFile(root); err != nil {
		return
	}

	b.Write([]byte("var files = []File{\n"))
	for _, f := range fs {
		b.Write([]byte("\tFile{\n\t\tName: \""))
		b.WriteString(path.Join("/", prefix, strings.TrimPrefix(f, root)))
		b.Write([]byte("\",\n\t\tContent: []byte(\""))
		if buf, err = ioutil.ReadFile(f);err != nil {
			return
		}
		if buf, err = Compress(buf); err != nil {
			return
		}
		b.WriteString(base64.RawStdEncoding.EncodeToString(buf))
		b.Write([]byte("\"),\n\t},\n"))
	}
	b.Write([]byte("}"))
	buffer = b.Bytes()
	return
}

func parseArgs() (root, prefix, output string, err error) {
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
	root, prefix, output, err := parseArgs()
	if err != nil {
		fmt.Println(err)
		return
	}
	buf, err := TarCompressAllFile(root, prefix)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = ioutil.WriteFile(output, buf, os.ModePerm); err != nil {
		fmt.Println(err)
	}
}