package fsdeser

import "testing"
import a "github.com/stretchr/testify/assert"

func Test_String(t *testing.T) {
	assert := a.New(t)

	fs := NewFs("dupastring.txt", &DeSerString{})

	err := fs.Store("kot")
	assert.NoError(err, "store shouldn't fail")

	ret, err := fs.Read()
	assert.NoError(err, "read shouldn't fail")
	assert.Equal("kot", ret, "should be same")
}

func Test_Int(t *testing.T) {
	assert := a.New(t)

	fs := NewFs("dupaint.txt", &DeSerInt{})

	err := fs.Store(666)
	assert.NoError(err, "store shouldn't fail")

	ret, err := fs.Read()
	assert.NoError(err, "read shouldn't fail")
	assert.Equal(666, ret, "should be same")

}
