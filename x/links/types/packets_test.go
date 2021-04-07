package types_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/desmos-labs/desmos/x/links/types"
	"github.com/stretchr/testify/require"
)

func TestIBCAccountConnectionPacketData_Validate(t *testing.T) {
	tests := []struct {
		name   string
		packet types.IBCAccountConnectionPacketData
		expErr error
	}{
		{
			name: "Valid IBCAccountConnectionPacketData",
			packet: types.NewIBCAccountConnectionPacketData(
				"cosmos",
				"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
				"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
			),
			expErr: nil,
		},
		{
			name: "Source Pubkey and Address Mismatch",
			packet: types.NewIBCAccountConnectionPacketData(
				"desmos",
				"desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
				"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
			),
			expErr: fmt.Errorf("source pubkey mismatch with source address"),
		}, {
			name: "Source Verify Error",
			packet: types.NewIBCAccountConnectionPacketData(
				"desmos",
				"desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk",
				"033162405bee8a826a3d4a62842f525f1e88f821a6225289b3d44c209be41c257b",
				"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
				"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
			),
			expErr: fmt.Errorf("source pubkey mismatch with source address"),
		}, {
			name: "Valid IBCAccountConnectionPacketData with Different Prefix",
			packet: types.NewIBCAccountConnectionPacketData(
				"desmos",
				"desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk",
				"02466b245623786131225676fbcf4eb5a32c835a8acc733a989af45b0cbbcc0e84",
				"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
				"28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
				"fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
			),
			expErr: nil,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.packet.Validate())
		})
	}
}

func Test_verify(t *testing.T) {
	tests := []struct {
		name   string
		msg    string // Hex-encoded msg
		sig    string // Hex-encoded sig
		pubKey string // Hex-encoded pubKey
		expRes bool
	}{
		{
			name:   "Valid Signature",
			msg:    "28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			sig:    "fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
			pubKey: "02b493a33f104de068e93d51ffe9929409a20635a68d0c2bc2b51d95e186e58f07",
			expRes: true,
		},
		{
			name:   "Invalid Signature",
			msg:    "28620f478ad11508ff4fbd01554f6dc4870e6d0ac656221774cabf9cef60951956324097b8642c0d09d23ab37bf0d6c1ea02816d92a0251acab42097a25e74b2",
			sig:    "fc0bc7dd041c736b8fa3bb6638fc003944b430aaa656d08b823836894338d30d5bb8c96e43d4c40d820acf2f6d03c8123df525c59eed114564b877ed1f7dd561",
			pubKey: "02466b245623786131225676fbcf4eb5a32c835a8acc733a989af45b0cbbcc0e84",
			expRes: false,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			msg, _ := hex.DecodeString(test.msg)
			sig, _ := hex.DecodeString(test.sig)
			pubKeyBs, _ := hex.DecodeString(test.pubKey)
			pubKey := &secp256k1.PubKey{Key: pubKeyBs}
			require.Equal(t, test.expRes, types.VerifySignature(msg, sig, pubKey))
		})
	}
}