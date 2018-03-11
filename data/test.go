// +build !notest

package data

import (
	"crypto/ecdsa"
	cryptorand "crypto/rand"
	"log"
	"math/rand"

	"github.com/ethereum/go-ethereum/common/number"
	"github.com/ethereum/go-ethereum/crypto"
	reform "gopkg.in/reform.v1"

	"github.com/privatix/dappctrl/util"
)

// TestData is a container for testing data items.
type TestData struct {
	Channel  string
	Password string
}

// These are functions for shortening testing boilerplate.

// NewTestDB creates a new database connection.
func NewTestDB(conf *DBConfig, logger *util.Logger) *reform.DB {
	db, err := NewDB(conf, logger)
	if err != nil {
		log.Fatalf("failed to open db connection: %s\n", err)
	}
	return db
}

// NewTestSubject returns new subject
func NewTestSubject() *Subject {
	priv, _ := ecdsa.GenerateKey(crypto.S256(), cryptorand.Reader)
	b := crypto.FromECDSA(priv)
	privB64 := FromBytes(b)
	priv, _ = crypto.ToECDSA(b)
	pub := FromBytes(
		crypto.FromECDSAPub(&priv.PublicKey))
	return &Subject{
		ID:         util.NewUUID(),
		PrivateKey: &privB64,
		PublicKey:  pub,
	}
}

// NewTestOffering returns new offering.
func NewTestOffering(agent *Subject) *Offering {
	return &Offering{
		ID:      util.NewUUID(),
		Agent:   agent.ID,
		Service: ServiceVPN,
		Supply:  1,
	}
}

// NewTestChannel returns new channel.
func NewTestChannel(agent, client *Subject, offering *Offering,
	balance, deposit int64, state string) *Channel {
	return &Channel{
		ID:             util.NewUUID(),
		Agent:          agent.ID,
		Client:         client.ID,
		Offering:       offering.ID,
		Block:          uint(rand.Intn(99999999)),
		State:          state,
		TotalDeposit:   FromBytes(number.Big(deposit).Bytes()),
		ClosedDeposit:  FromBytes(number.Big(0).Bytes()),
		ReceiptBalance: FromBytes(number.Big(balance).Bytes()),
	}
}
