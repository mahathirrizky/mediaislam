package ustadz

import "time"

type UstadzTable struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
