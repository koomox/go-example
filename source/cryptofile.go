package main

import (
	"github.com/koomox/ext"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func main() {
	loadLog()
	src, dst, secret, action, big, err := parseArgs()
	switch action {
	case "encode":
		switch big {
		case true:
			if err = ext.EncryptBigFile(src, dst, []byte(secret)); err != nil {
				log.Errorf("EncryptBigFile Err:%v", err.Error())
				return
			}
		default:
			if err = ext.EncryptFile(src, dst, []byte(secret)); err != nil {
				log.Errorf("EncryptFile Err:%v", err.Error())
				return
			}
		}

		log.Infof("EncryptFile src(\"%v\") => dst(\"%v\")", src, dst)
	case "decode":
		switch big {
		case true:
			if err = ext.DecryptBigFile(src, dst, []byte(secret)); err != nil {
				log.Errorf("EncryptBigFile Err:%v", err.Error())
				return
			}
		default:
			if err = ext.DecryptFile(src, dst, []byte(secret)); err != nil {
				log.Errorf("EncryptFile Err:%v", err.Error())
				return
			}
		}
		log.Infof("DecryptFile src(\"%v\") => dst(\"%v\")", src, dst)
	}
}

func loadLog() {
	customFormatter := &log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	log.SetFormatter(customFormatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func parseArgs() (src, dst, secret, action string, big bool, err error) {
	big = false
	for _, arg := range os.Args {
		op := strings.Split(arg, "=")
		sc := op[0]
		switch len(op) {
		case 1:
			switch sc {
			case "--encode":
				action = "encode"
			case "--decode":
				action = "decode"
			case "--big":
				big = true
			}
		case 2:
			switch sc {
			case "--in":
				src = op[1]
			case "--out":
				dst = op[1]
			case "--secret":
				secret = op[1]
			}
		}
	}

	switch src {
	case "":
		log.Errorf("srcFile Not Null!")
		return
	}

	switch dst {
	case "":
		if dst, err = ext.GeneratorRawFileNameUnique(""); err != nil {
			log.Errorf("Generator Unique FileName Err:%v", err.Error())
			return
		}
	}

	switch secret {
	case "":
		if secret, err = ext.RandomString(32); err != nil {
			log.Errorf("Generator Secret Err:%v", err.Error())
			return
		}
		log.Infof("Generator Secret(\"%v\")", secret)
	}

	switch action {
	case "":
		log.Errorf("action Not Null!")
		return
	}

	return
}