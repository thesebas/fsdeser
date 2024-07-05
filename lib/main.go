package fsdeser

import (
	"os"
	"strconv"
)

type FS[T any] struct {
	filename string
	deser    DeSer[T]
}

func NewFs[S DeSer[T], T any](filename string, deser S) *FS[T] {
	return &FS[T]{
		filename,
		deser,
	}
}

func (fs *FS[T]) Store(data T) error {
	b := fs.deser.Serialize(data)
	return os.WriteFile(fs.filename, b, 0666)
}

func (fs *FS[T]) Read() (T, error) {
	b, err := os.ReadFile(fs.filename)
	if err != nil {
		return fs.deser.Empty(), err
	}
	d, err := fs.deser.Deserialize(b)
	if err != nil {
		return fs.deser.Empty(), err
	}
	return d, nil
}

// region DeSerString

type DeSerString struct {
}

func (d *DeSerString) Deserialize(b []byte) (string, error) {
	return string(b), nil
}

func (d *DeSerString) Serialize(s string) []byte {
	return []byte(s)
}
func (d *DeSerString) Empty() string {
	return ""
}

// endregion
// region DeSerInt

type DeSerInt struct {
}

func (d *DeSerInt) Deserialize(b []byte) (int, error) {
	return strconv.Atoi(string(b))
}

func (d *DeSerInt) Serialize(s int) []byte {
	return []byte(strconv.Itoa(s))
}
func (d *DeSerInt) Empty() int {
	return 0
}

// endregion

type DeSer[T any] interface {
	Serialize(T) []byte
	Deserialize([]byte) (T, error)
	Empty() T
}
