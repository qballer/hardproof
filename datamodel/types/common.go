package types

import (
	"math/big"
	"crypto/ecdsa"
)

type Key byte[32]
type Signature byte[32]
type Address byte[32]

type SignedPublicKey struct {
	Key
	Signature
}

type Keypair interface {
	PrivateKey Key
	PublicKey Key
}

type Attestation interface {

}

type BlockID interface {

}

type TxID interface {

}


type Payload interface {

}

type SignedPayload interface {

}
