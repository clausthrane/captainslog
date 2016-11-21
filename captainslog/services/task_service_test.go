package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil("fooo", "expecting err")
}

