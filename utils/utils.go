package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jack-zh/ztodo/zterminal/gopass"
	"github.com/jack-zh/ztodo/zterminal/rgbterm"
)

func RemoveSlice(slice []interface{}, start, end int) []interface{} {
	result := make([]interface{}, len(slice)-(end-start))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end:])
	return result
}

func Credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	password := string(gopass.GetPasswd())

	return strings.TrimSpace(username), strings.TrimSpace(password)
}

func CredentialsRetype() (string, string, string) {
	word := "=)"
	var r, g, b uint8
	r, g, b = 43, 43, 43
	coloredWord := rgbterm.BgString(word, r, g, b)
	fmt.Println("Oh!", coloredWord, "hello!")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	password := string(gopass.GetPasswd())

	fmt.Print("Enter Retype Password: ")
	retypepassword := string(gopass.GetPasswd())

	return strings.TrimSpace(username), strings.TrimSpace(password), strings.TrimSpace(retypepassword)
}
