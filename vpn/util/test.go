// +build !notest

package util

import (
	"testing"

	"gopkg.in/reform.v1"
)

// SetChannelState assigns state to a given state channel.
func SetChannelState(t *testing.T, db *reform.DB, ch, state string) {
	if _, err := db.Exec(`UPDATE channels
		                 SET state = $1
		               WHERE id = $2`, state, ch); err != nil {
		t.Fatalf("failed to set channel state: %s", err)
	}
}
