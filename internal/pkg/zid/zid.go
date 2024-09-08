package zid

import "github.com/costa92/krm/pkg/id"

const defaultABC = "abcdefghijklmnopqrstuvwxyz1234567890"

type ZID string

const (
	User  ZID = "user"
	Order ZID = "order"
)

func (zid ZID) String() string {
	return string(zid)
}

func (zid ZID) New(i uint64) string {
	str := id.NewCode(
		i,
		id.WithCodeChars([]rune(defaultABC)),
		id.WithCodeL(6),
		id.WithCodeSalt(Salt()),
	)
	return zid.String() + "-" + str
}
