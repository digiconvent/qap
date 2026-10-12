package layer3

import "time"

func StartCleanup(store *Store, ticks <-chan time.Time, now func() time.Time) {
	go func() {
		for range ticks {
			store.RemoveExpired(now().Unix())
		}
	}()
}
