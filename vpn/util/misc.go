package util

import (
	"gopkg.in/reform.v1"

	"github.com/privatix/dappctrl/data"
)

// FindCurrentSession finds a current session for a given state channel.
func FindCurrentSession(db *reform.DB, ch string) (string, error) {
	var sess string
	if err := db.QueryRow(`
		SELECT id
		  FROM sessions
		 WHERE channel = $1 AND stopped IS NULL
		 ORDER BY started DESC
		 LIMIT 1`, ch).Scan(&sess); err != nil {
		return "", err
	}

	return sess, nil
}

// ChannelOpen checks whether a given state channel is open. For VPN service
// that means it allows to connect and to consume traffic.
func ChannelOpen(db *reform.DB, ch string) (bool, error) {
	var open bool
	if err := db.QueryRow(`
		SELECT state = $1
		  FROM channels
		 WHERE id = $2`, data.ChannelOpen, ch).Scan(&open); err != nil {
		return false, err
	}

	return open, nil
}
