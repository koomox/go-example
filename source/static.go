package main

import (
	"bytes"
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
	ContentType string
	Content []byte
}

type Item struct {
	root string
	filter []string
	prefix string
}

type Encoding struct {
	buffer bytes.Buffer
	item []*Item
}

func NewEncoding() *Encoding {
	return &Encoding{}
}

func DeCompress(data []byte) (buf []byte, err error) {
	if data, err = base64.RawStdEncoding.DecodeString(string(data)); err != nil {
		return
	}

	return ext.NewEncoding().DeCompress(data)
}

func (this *Encoding)CompressAllFile() (buffer []byte, err error) {
	var (
		b bytes.Buffer
		buf []byte
	)

	b.Write([]byte("var files = []File{\n"))
	for _, item := range this.item {
		if buf, err = item.Compress(); err != nil {
			return
		}
		b.Write(buf)
	}
	b.Write([]byte("}"))
	buffer = b.Bytes()
	return
}

func (this *Encoding)New() (*Item) {
	item := &Item{}
	this.item = append(this.item, item)
	return item
}

func (this *Item)Compress() (buffer []byte, err error){
	var (
		fs []string
		b bytes.Buffer
		buf []byte
	)
	if this.root == "" {
		if this.root, err = ext.GetCurrentDirectory(); err != nil {
			return
		}
	}
	this.root = path.Join(this.root, "")
	if fs, err = ext.GetCustomDirectoryAllFile(this.root); err != nil {
		return
	}

	for _, f := range fs {
		filter := false
		for _, v := range this.filter {
			if strings.EqualFold(f, v) || strings.HasPrefix(f, v) {
				filter = true
				break
			}
		}
		if filter {
			continue
		}
		b.Write([]byte("\tFile{\n\t\tName: \""))
		b.WriteString(path.Join(this.prefix, strings.TrimPrefix(f, this.root)))
		b.Write([]byte("\",\n\t\tContentType: \""))
		switch path.Ext(f) {
		case ".js":
			b.WriteString("text/javascript; charset=utf-8")
		case ".css":
			b.WriteString("text/css; charset=utf-8")
		case ".html", ".htm", ".php":
			b.WriteString("text/html; charset=utf-8")
		case ".jpg", "jpeg":
			b.WriteString("image/jpeg; charset=utf-8")
		case ".gif":
			b.WriteString("image/gif; charset=utf-8")
		case ".png":
			b.WriteString("image/png; charset=utf-8")
		case ".otf":
			b.WriteString("font/otf; charset=utf-8")
		case ".ttf":
			b.WriteString("font/ttf; charset=utf-8")
		case ".woff":
			b.WriteString("font/woff; charset=utf-8")
		case ".woff2":
			b.WriteString("font/woff2; charset=utf-8")
		default:
			b.WriteString("text/plain; charset=utf-8")
		}
		b.Write([]byte("\",\n\t\tContent: []byte(\""))
		if buf, err = ioutil.ReadFile(f);err != nil {
			return
		}
		if buf, err = ext.NewEncoding().Compress(buf); err != nil {
			return
		}
		b.WriteString(base64.RawStdEncoding.EncodeToString(buf))
		b.Write([]byte("\"),\n\t},\n"))
	}
	buffer = b.Bytes()
	return
}

func (this *Item)Filter(elem ...string) (*Item){
	for i, e := range elem {
		if e != "" {
			this.filter = append(this.filter, elem[i])
		}
	}

	return this
}

func (this *Item)Root(root string) (*Item){
	root = strings.Replace(root, "\\", "/", -1)
	this.root = root

	return this
}

func (this *Item)Prefix(prefix string) (*Item){
	prefix = strings.Replace(prefix, "\\", "/", -1)
	this.prefix = path.Join("/", prefix)
	return this
}

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
	enc := NewEncoding()
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