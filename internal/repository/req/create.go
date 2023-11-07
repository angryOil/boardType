package req

import "time"

type Create struct {
	Name        string
	Description string
	CafeId      int
	CreateBy    int
	CreatedAt   time.Time
}
