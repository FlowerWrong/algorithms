package linearlist

import (
	"testing"

	"."
	"github.com/stretchr/testify/assert"
)

// go test *_test.go -test.bench=".*"

func TestSqList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	l := linearlist.NewSqList(100)
	assert.Equal(t, true, l.IsEmpty(), "they should be equal")
	assert.Equal(t, false, l.IsFull(), "they should be equal")

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Insert(2, 4)
	l.Insert(3, 5)
	assert.Equal(t, 5, l.Len(), "they should be equal")

	e, err := l.Get(2)
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, 4, int(e), "they should be equal")

	assert.Equal(t, []linearlist.Elem{1, 4, 5, 2, 3}, l.Data(), "they should be equal")

	l.Del(2)
	assert.Equal(t, 4, l.Len(), "they should be equal")

	e, err = l.Get(2)
	assert.Equal(t, nil, err, "they should be equal")
	assert.Equal(t, 5, int(e), "they should be equal")

	l.Clear()
	assert.Equal(t, true, l.IsEmpty(), "they should be equal")
}
