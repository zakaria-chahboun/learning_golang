package bill

import "time"

type Product struct {
	ID    int
	Title string
	Price float64
}

type Invoice struct {
	ID      int
	Product Product
	Date    time.Time
}

func (i Invoice) IsOutdated() bool {
	out := time.Now().Add(time.Hour * 24 * 5)
	if i.Date.Unix() >= out.Unix() {
		return true
	}
	return false
}
