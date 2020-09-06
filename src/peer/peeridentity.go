package peer

import (
	"errors"
)

type PeerIdentity struct {
	peeridentitytype int
	publickey        []byte
	publickeyhash    []byte
	shortid          string
}

func CreateIdentity(peeridentitytype int) (*PeerIdentity, error) {
	switch peeridentitytype {
	case 1:
		break
	default:
		return nil, errors.New("only peeridentitype 1 is supported")
	}

	var identity *PeerIdentity
	identity = new(PeerIdentity)
	// implementation
	return identity, nil
}
