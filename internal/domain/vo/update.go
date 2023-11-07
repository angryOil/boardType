package vo

import "time"

type Update struct {
	Id          int
	CreateBy    int
	CafeId      int
	Name        string
	Description string
	CreatedAt   time.Time
}
