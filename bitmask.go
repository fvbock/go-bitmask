package bitmask

import (
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

func (b *bitmask) IsSet(value int) {
	return b.values.HasMember(value)
}

func (b *bitmask) setValues() {
	i := 1
	for i <= value {
		if i&f == i {
			b.values.Add(f)
		}
		i = i * 2
	}

	return
}

func (b *bitmask) String() string {
	return fmt.Springf("%v", b.values.Members())
}
