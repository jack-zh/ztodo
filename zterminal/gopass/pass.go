package gopass

import (
	"os"
)

func getPasswd(masked bool) []byte {
	var pass, bs, mask []byte
	if masked {
		bs = []byte("\b \b")
		mask = []byte("*")
	}
	for {
		if v := getch(); v == 127 || v == 8 {
			if l := len(pass); l > 0 {
				pass = pass[:l-1]
				os.Stdout.Write(bs)
			}
		} else if v == 13 || v == 10 {
			break
		} else {
			pass = append(pass, v)
			os.Stdout.Write(mask)
		}
	}
	println()
	return pass
}

func GetPasswd() []byte {
	return getPasswd(false)
}

func GetPasswdMasked() []byte {
	return getPasswd(true)
}
