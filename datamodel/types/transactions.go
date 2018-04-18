package types

import (
	"math/big"
)

type Transaction interface {
	Sender() Address
	SenderSignature Signature
}

type TransactionSet interface {

}

type TxTransfer interface {
	From Address
	To Address
	Amount Int



}

type TxValidate interface {

}

type TxApproveMiningKey interface {
	Attestation
}

type TxApproveHardware interface {
	HardwareIdentifier() string
	Votes() []Signature
}

type TxRevokeHardware interface {
	HardwareIdentifier() string
	Votes() []Signature
}

type TxCreateCurrency interface {
	TokenAddress() Address
	InitialSupply() Int
	Precision int32
}

type TxCreateMiningToken interface {
	PublicKey() SignedPublicKey
}

