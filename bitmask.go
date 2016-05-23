package bitmask

import (
	"fmt"

	"github.com/fvbock/uds-go/set"
)

type bitmask struct {
	int
	values *set.IntSet
}

func Bitmask(value int) (b *bitmask) {
	b = &bitmask{
		value,
		set.NewIntSet(),
	}

	b.setValues()

	return
}

func (b *bitmask) IsSet(value int) bool {
	return b.values.HasMember(value)
}

func (b *bitmask) setValues() {
	i := 1
	for i <= b.int {
		if i&b.int == i {
			b.values.Add(i)
		}
		i = i * 2
	}

	return
}

func (b *bitmask) String() string {
	return fmt.Sprintf("%v", b.values.Members())
}
