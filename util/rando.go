package util

import (
	"math/rand"
	"strings"
	"time"
)

const chars = "abcdefghijklmnopqrstuvwxyz123456789"
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOQRZTUVWXYZ"

var names = []string{"John", "Paul", "Marcus", "Vanessa", "Paulina", "Xavier", "Richard", "Irene"}
var roles = []string{"Admin", "Associate", "Member"}
var status = []string{"Open", "Closed", "In Progress", "In Progress"}
var severityLevel = []string{"Low", "Moderate", "High", "Urgent"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomNumber(n int64) int64 {
	return rand.Int63n(n)
}

func RandomName() string {
	n := len(names)

	return names[rand.Intn(n)]
}

func RandomChars(n int) string {
	var sb strings.Builder

	a := len(chars)
	for i := 0; i < n; i++ {
		c := chars[rand.Intn(a)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomRole() string {
	n := len(roles)

	return roles[rand.Intn(n)]
}

func RandomCompany(n int) string {
	var sb strings.Builder

	a := len(alphabet)
	for i := 0; i < n; i++ {
		c := chars[rand.Intn(a)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomStatus() string {
	rStatus := len(status)

	return status[rand.Intn(rStatus)]
}

func RandomSeverityLevel() string {

	rSLevel := len(severityLevel)

	return severityLevel[rand.Intn(rSLevel)]
}

func RandomBool() bool {

	rand := rand.Intn(2)

	if rand == 1 {
		return false
	}
	return true
}
