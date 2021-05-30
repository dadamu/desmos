package keeper_test

import (
	"encoding/hex"
	"time"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	"github.com/desmos-labs/desmos/testutil/ibctesting"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestOnRecvPacket() {
	var (
		channelA, channelB ibctesting.TestChannel
		srcAddr            string
		destAddr           string
		srcSigHex          string
		destSigHex         string
	)

	tests := []struct {
		name         string
		malleate     func()
		storeProfile func()
		doubleStore  bool
		expPass      bool
	}{
		{
			name: "Create link from source chain successfully",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()

				srcSig, _ := suite.chainA.PrivKey.Sign([]byte(srcAddr))
				srcSigHex = hex.EncodeToString(srcSig)

				destAddr = suite.chainB.Account.GetAddress().String()
				dstSig, _ := suite.chainB.PrivKey.Sign([]byte(destAddr))
				destSigHex = hex.EncodeToString(dstSig)
			},
			storeProfile: func() {
				addr := suite.chainB.Account.GetAddress()
				baseAcc := authtypes.NewBaseAccountWithAddress(addr)
				baseAcc.SetPubKey(suite.chainB.Account.GetPubKey())

				profile, err := types.NewProfile(
					"dtag",
					"test-user",
					"biography",
					types.NewPictures(
						"https://shorturl.at/adnX3",
						"https://shorturl.at/cgpyF",
					),
					time.Time{},
					baseAcc,
				)
				suite.Require().NoError(err)
				err = suite.chainB.App.ProfileKeeper.StoreProfile(suite.chainB.GetContext(), profile)
				suite.Require().NoError(err)
			},
			expPass: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupIBCTest()
			test.malleate()

			packetData := types.NewLinkChainAccountPacketData(
				srcAddr,
				types.NewProof(
					suite.chainA.Account.GetPubKey(),
					srcSigHex,
					srcAddr,
				),
				types.NewChainConfig(
					"test",
					"cosmos",
				),
				destAddr,
				types.NewProof(
					suite.chainB.Account.GetPubKey(),
					destSigHex,
					destAddr,
				),
			)
			bz, _ := packetData.GetBytes()
			packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)

			test.storeProfile()
			_, err := suite.chainB.App.ProfileKeeper.OnRecvPacket(
				suite.chainB.GetContext(),
				packet,
				packetData,
			)
			if test.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestOnAcknowledgementPacket() {
	var (
		channelA, channelB ibctesting.TestChannel
		ack                channeltypes.Acknowledgement
	)
	tests := []struct {
		name     string
		malleate func()
		expError error
	}{
		{
			name: "Receive success ack returns no error",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				packetAck := types.LinkChainAccountPacketAck{SourceAddress: suite.chainA.Account.GetAddress().String()}
				bz, _ := packetAck.Marshal()
				ack = channeltypes.NewResultAcknowledgement(bz)
			},
			expError: nil,
		},
	}

	for _, test := range tests {
		test := test

		suite.Run(test.name, func() {
			suite.SetupIBCTest()
			test.malleate()

			srcAddr := suite.chainA.Account.GetAddress().String()
			destAddr := suite.chainA.Account.GetAddress().String()

			srcSig, err := suite.chainA.PrivKey.Sign([]byte(srcAddr))
			suite.Require().NoError(err)
			srcSigHex := hex.EncodeToString(srcSig)
			destSig, err := suite.chainB.PrivKey.Sign([]byte(destAddr))
			suite.Require().NoError(err)
			destSigHex := hex.EncodeToString(destSig)

			data := types.NewLinkChainAccountPacketData(
				srcAddr,
				types.NewProof(
					suite.chainA.Account.GetPubKey(),
					srcSigHex,
					srcAddr,
				),
				types.NewChainConfig(
					"test",
					"cosmos",
				),
				destAddr,
				types.NewProof(
					suite.chainB.Account.GetPubKey(),
					destSigHex,
					destAddr,
				),
			)
			bz, _ := data.GetBytes()
			packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
			err = suite.chainA.App.ProfileKeeper.OnAcknowledgementPacket(suite.chainA.GetContext(), packet, data, ack)
			if test.expError == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
