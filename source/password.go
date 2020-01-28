package main


import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
	"strconv"
)

const (
	randomComplete = "`~^0OolI\"'/\\|"
	randomLetter   = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	randomNumber   = "123456789"
	randomSpecial  = "~!@#$%^&*_+:?`-=;,."
)

func generatorRandomStr(length int, complete, noNumber, noSpecial bool) (string, error) {
	var randomPool = randomLetter

	if complete {
		if noNumber || noSpecial {
			return "", fmt.Errorf("Cannot use `complete` flag with `no-number` and `no-special`.")
		}
		randomPool += randomNumber
		randomPool += randomSpecial
		randomPool += randomComplete
	} else {
		if !noNumber {
			randomPool += randomNumber
		}

		if !noSpecial {
			randomPool += randomSpecial
		}
	}

	randstr := make([]byte, length) // Random string to return
	charlen := big.NewInt(int64(len(randomPool)))
	for i := 0; i < length; i++ {
		b, err := rand.Int(rand.Reader, charlen)
		if err != nil {
			return "", fmt.Errorf("RandString Generator Err:%v", err.Error())
		}
		r := int(b.Int64())
		randstr[i] = randomPool[r]
	}
	return string(randstr), nil
}

func RandomString(length int) (result string, err error) {
	for i := 0; i < 3; i++ {
		if result, err = generatorRandomStr(length, false, false, true); err != nil {
			continue
		}
		return
	}
	return
}

func main() {
	var (
		length int
		number int
		symbol bool
		result string
		err error
	)
	length, number, symbol = parseArg()
	for i:=0; i < number; i++ {
		if symbol {
			result, err = generatorRandomStr(length, false, false, false)
		} else {
			result, err = generatorRandomStr(length, false, false, true)
		}

		if err != nil {
			fmt.Println(err.Error())
			break
		}

		if i+1 != number {
			fmt.Printf("\"%v\",\n", result)
		} else {
			fmt.Printf("\"%v\"", result)
		}
	}
}


func parseArg() (length, number int, symbol bool) {
	var (
		lengthString string
		numberString string
		err error
	)

	for _, arg := range os.Args {
		option := strings.Split(arg, "=")
		item := option[0]
		switch len(option) {
		case 1:
			switch strings.ToLower(item) {
			case "--s", "--sym", "--symbol":
				symbol = true
			}
		case 2:
			switch strings.ToLower(item) {
			case "--l", "--len", "--length":
				lengthString = option[1]
			case "--n", "--num", "--number":
				numberString = option[1]
			}
		}
	}

	if lengthString == "" {
		length = 32
	} else {
		if length, err = strconv.Atoi(lengthString); err != nil {
			length = 32
		}
	}

	if numberString == "" {
		number = 1
	} else {
		if number, err = strconv.Atoi(numberString); err != nil {
			number = 1
		}
	}

	return
}