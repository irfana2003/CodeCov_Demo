package cal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T){
	assert := assert.New(t)
	result, err := Add(1,2)
	assert.Equal(result, 3)
	assert.NoError(err)
}

func TestSub(t *testing.T){
	assert := assert.New(t)
	result, err := Sub(1,2)
	assert.Equal(result, -1)
	assert.NoError(err)
}

func TestMul(t *testing.T){
	assert := assert.New(t)
	result, err := Mul(1,2)
	assert.Equal(result, 2)
	assert.NoError(err)
}

func TestDiv(t *testing.T){
	assert := assert.New(t)
	result, err := Div(4,2)
	assert.Equal(result, 2)
	assert.NoError(err)
}