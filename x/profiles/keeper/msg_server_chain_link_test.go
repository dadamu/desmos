package keeper_test

import (
	"encoding/hex"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_handleMsgLinkChainAccount() {
	var existentLinks []types.ChainLink
	var existentProfiles []*types.Profile

	// Generate source and destination key
	srcPriv := secp256k1.GenPrivKey()
	destPriv := secp256k1.GenPrivKey()
	srcPubKey := srcPriv.PubKey()
	destPubKey := destPriv.PubKey()

	// Get bech32 encoded addresses
	srcAddr, err := bech32.ConvertAndEncode("cosmos", srcPubKey.Address().Bytes())
	suite.Require().NoError(err)
	destAddr, err := bech32.ConvertAndEncode("cosmos", destPubKey.Address().Bytes())
	suite.Require().NoError(err)

	// Get signature by signing with keys
	srcSig, err := srcPriv.Sign([]byte(srcAddr))
	suite.Require().NoError(err)
	destSig, err := destPriv.Sign([]byte(destAddr))
	suite.Require().NoError(err)

	srcSigHex := hex.EncodeToString(srcSig)
	destSigHex := hex.EncodeToString(destSig)

	blockTime := suite.testData.profile.CreationDate

	tests := []struct {
		name           string
		malleate       func()
		msg            *types.MsgLinkChainAccount
		shouldErr      bool
		expEvents      sdk.Events
		expStoredLinks []types.ChainLink
	}{
		{
			name: "Create link successfully",
			malleate: func() {
				srcAccAddr, err := sdk.AccAddressFromBech32(srcAddr)
				suite.Require().NoError(err)

				srcBaseAcc := authtypes.NewBaseAccountWithAddress(srcAccAddr)
				srcBaseAcc.SetPubKey(srcPubKey)
				suite.ak.SetAccount(suite.ctx, srcBaseAcc)

				destAccAddr, err := sdk.AccAddressFromBech32(destAddr)
				suite.Require().NoError(err)
				destBaseAcc := authtypes.NewBaseAccountWithAddress(destAccAddr)
				destBaseAcc.SetPubKey(destPubKey)
				suite.ak.SetAccount(suite.ctx, destBaseAcc)

				existentProfiles = []*types.Profile{
					suite.CheckProfileNoError(types.NewProfile(
						"custom_dtag",
						"my-nickname",
						"my-bio",
						types.NewPictures(
							"https://test.com/profile-picture",
							"https://test.com/cover-pic",
						),
						suite.testData.profile.CreationDate,
						destBaseAcc,
					)),
				}

				existentLinks = nil
			},
			msg: types.NewMsgLinkChainAccount(
				srcAddr,
				types.NewProof(srcPubKey, srcSigHex, srcAddr),
				types.NewChainConfig("cosmos", sdk.GetConfig().GetBech32AccountAddrPrefix()),
				destAddr,
				types.NewProof(destPubKey, destSigHex, destAddr),
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeLinkChainAccount,
					sdk.NewAttribute(types.AttributeChainLinkAccountTarget, srcAddr),
					sdk.NewAttribute(types.AttributeChainLinkSourceChainName, "cosmos"),
					sdk.NewAttribute(types.AttributeChainLinkAccountOwner, destAddr),
					sdk.NewAttribute(types.AttributeChainLinkCreated, blockTime.Format(time.RFC3339Nano)),
				),
			},
			expStoredLinks: []types.ChainLink{
				types.NewChainLink(
					srcAddr,
					types.NewProof(srcPubKey, srcSigHex, srcAddr),
					types.NewChainConfig("cosmos", sdk.GetConfig().GetBech32AccountAddrPrefix()),
					blockTime,
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			test.malleate()
			suite.ctx = suite.ctx.WithBlockTime(blockTime)
		})
		for _, acc := range existentProfiles {
			err := suite.k.StoreProfile(suite.ctx, acc)
			suite.Require().NoError(err)
		}

		for _, link := range existentLinks {
			err := suite.k.StoreChainLink(suite.ctx, link, existentProfiles[0].GetAddress().String())
			suite.Require().NoError(err)
		}

		server := keeper.NewMsgServerImpl(suite.k)
		_, err := server.LinkChainAccount(sdk.WrapSDKContext(suite.ctx), test.msg)
		suite.Require().NoError(err)
		suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

		if test.shouldErr {
			suite.Require().Error(err)
		} else {
			suite.Require().NoError(err)

			stored := suite.k.GetAllChainsLinks(suite.ctx)
			suite.Require().NotEmpty(stored)

			suite.Require().True(checkAllLinkContains(test.expStoredLinks, stored))

			profile, found, err := suite.k.GetProfile(suite.ctx, destAddr)
			suite.Require().NoError(err)
			suite.Require().True(found)

			suite.Require().NotEmpty(profile.ChainsLinks)
			suite.Require().True(checkAllLinkContains(test.expStoredLinks, profile.ChainsLinks))
		}
	}
}

func checkAllLinkContains(expStoredLinks []types.ChainLink, storedLinks []types.ChainLink) bool {
	allContain := false
	for _, link := range storedLinks {
		allContain = false
		for _, expLink := range expStoredLinks {
			if link.Equal(expLink) {
				allContain = true
				break
			}
		}
	}
	return allContain
}

func (suite *KeeperTestSuite) Test_handleMsgUnlinkChainAccount() {
	// Generate source and destination key
	srcPriv := secp256k1.GenPrivKey()
	srcPubKey := srcPriv.PubKey()

	// Get bech32 encoded addresses
	srcAddr, err := bech32.ConvertAndEncode("cosmos", srcPubKey.Address().Bytes())
	suite.Require().NoError(err)

	// Get signature by signing with keys
	srcSig, err := srcPriv.Sign([]byte(srcAddr))
	suite.Require().NoError(err)

	srcSigHex := hex.EncodeToString(srcSig)

	link := types.NewChainLink(
		srcAddr,
		types.NewProof(srcPubKey, srcSigHex, srcAddr),
		types.NewChainConfig("cosmos", "cosmos"),
		time.Time{},
	)

	validProfile := *suite.testData.profile
	validProfile.ChainsLinks = []types.ChainLink{
		link,
	}

	tests := []struct {
		name             string
		msg              *types.MsgUnlinkChainAccount
		shouldErr        bool
		expEvents        sdk.Events
		existentProfiles []*types.Profile
		existentLinks    []types.ChainLink
		expStoredLinks   []types.ChainLink
	}{
		{
			name:      "No error message",
			msg:       types.NewMsgUnlinkChainAccount(suite.testData.user, "cosmos", srcAddr),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeUnlinkChainAccount,
					sdk.NewAttribute(types.AttributeChainLinkAccountTarget, srcAddr),
					sdk.NewAttribute(types.AttributeChainLinkSourceChainName, "cosmos"),
					sdk.NewAttribute(types.AttributeChainLinkAccountOwner, suite.testData.user),
				),
			},
			existentProfiles: []*types.Profile{&validProfile},
			expStoredLinks:   nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			for _, acc := range test.existentProfiles {
				err := suite.k.StoreProfile(suite.ctx, acc)
				suite.Require().NoError(err)
				for _, link := range acc.ChainsLinks {
					err := suite.k.StoreChainLink(suite.ctx, link, acc.GetAddress().String())
					suite.Require().NoError(err)
				}
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.UnlinkChainAccount(sdk.WrapSDKContext(suite.ctx), test.msg)
			suite.Require().NoError(err)
			suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				stored := suite.k.GetAllChainsLinks(suite.ctx)
				suite.Require().Equal(test.expStoredLinks, stored)

				profile, found, err := suite.k.GetProfile(suite.ctx, suite.testData.user)
				suite.Require().NoError(err)
				suite.Require().True(found)
				suite.Require().Empty(profile.ChainsLinks)
			}
		})
	}
}
