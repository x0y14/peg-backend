package util

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ParseBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return b
}

func CreateDirectChatId(userId1 string, userId2 string) string {
	mem := []string{userId1, userId2}
	sort.Strings(mem)
	return fmt.Sprintf("di|%s", strings.Join(mem, "."))
}
