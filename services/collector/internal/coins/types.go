package coins

import "time"

// TrackedCoin stores the info for a tracked coin. Row in DB tracked_coins table.
type TrackedCoin struct {
	ID        int // primary key
	CmCID     int // Coinmarketcap ID
	Symbol    string
	Name      string
	Enableed  bool      // coin istracked or not
	CreatedAt time.Time // initial creation date in DB table
}
