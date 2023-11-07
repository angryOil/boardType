package req

import "time"

type Update struct {
	Id          int
	Name        string
	Description string
	CafeId      int
	MemberId    int
	CreatedAt   time.Time
}
