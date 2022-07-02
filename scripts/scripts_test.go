package scripts

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetAdminClaim(t *testing.T) {
	err := SetAdminClaim("739y5slHFgc7CBB4A1jhETzRYyt1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateAdminToken(t *testing.T) {
	token, err := GenerateAdminToken()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, token, "")
	fmt.Println(token)
}
