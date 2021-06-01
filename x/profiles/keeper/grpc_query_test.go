package keeper_test

import (
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) codeToAny(profile *types.Profile) *codectypes.Any {
	accountAny, err := codectypes.NewAnyWithValue(profile)
	suite.Require().NoError(err)
	return accountAny
}

func (suite *KeeperTestSuite) Test_Profile() {
	usecases := []struct {
		name           string
		storedProfiles []*types.Profile
		req            *types.QueryProfileRequest
		shouldErr      bool
		expResponse    *types.QueryProfileResponse
	}{
		{
			name:      "empty user returns error",
			req:       types.NewQueryProfileRequest(""),
			shouldErr: true,
		},
		{
			name:      "non existing DTag returns error",
			req:       types.NewQueryProfileRequest("invalid-dtag"),
			shouldErr: true,
		},
		{
			name:        "profile not found",
			req:         types.NewQueryProfileRequest("cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa"),
			shouldErr:   false,
			expResponse: &types.QueryProfileResponse{Profile: nil},
		},
		{
			name: "found profile - using dtag",
			storedProfiles: []*types.Profile{
				suite.testData.profile,
			},
			req:       types.NewQueryProfileRequest(suite.testData.profile.DTag),
			shouldErr: false,
			expResponse: &types.QueryProfileResponse{
				Profile: suite.codeToAny(suite.testData.profile),
			},
		},
		{
			name: "found profile - using address",
			storedProfiles: []*types.Profile{
				suite.testData.profile,
			},
			req:       types.NewQueryProfileRequest(suite.testData.profile.GetAddress().String()),
			shouldErr: false,
			expResponse: &types.QueryProfileResponse{
				Profile: suite.codeToAny(suite.testData.profile),
			},
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			suite.SetupTest()

			for _, profile := range uc.storedProfiles {
				suite.Require().NoError(suite.k.StoreProfile(suite.ctx, profile))
			}

			res, err := suite.k.Profile(sdk.WrapSDKContext(suite.ctx), uc.req)

			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				suite.Require().Equal(uc.expResponse, res)

				if uc.expResponse.Profile != nil {
					// Make sure the cached value is not nil (this is to grant that UnpackInterfaces work properly)
					suite.Require().NotNil(res.Profile.GetCachedValue())
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_DTagTransfers() {
	usecases := []struct {
		name           string
		storedRequests []types.DTagTransferRequest
		req            *types.QueryDTagTransfersRequest
		shouldErr      bool
		expResponse    *types.QueryDTagTransfersResponse
	}{
		{
			name:      "invalid user",
			req:       types.NewQueryDTagTransfersRequest("invalid-address"),
			shouldErr: true,
		},
		{
			name: "valid request",
			storedRequests: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
				),
				types.NewDTagTransferRequest(
					"dtag-2",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				),
			},
			req:       types.NewQueryDTagTransfersRequest("cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa"),
			shouldErr: false,
			expResponse: &types.QueryDTagTransfersResponse{
				Requests: []types.DTagTransferRequest{
					types.NewDTagTransferRequest(
						"dtag",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					),
				},
			},
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			suite.SetupTest()

			for _, req := range uc.storedRequests {
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(suite.ctx, req))
			}

			res, err := suite.k.DTagTransfers(sdk.WrapSDKContext(suite.ctx), uc.req)

			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				suite.Require().Equal(uc.expResponse, res)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_Params() {
	suite.k.SetParams(suite.ctx, types.DefaultParams())

	res, err := suite.k.Params(sdk.WrapSDKContext(suite.ctx), &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	suite.Require().NotNil(res)

	suite.Require().Equal(types.DefaultParams(), res.Params)
}

func (suite *KeeperTestSuite) Test_ChainsLinks() {
	suite.SetupTest()
	profile := suite.testData.profile

	suite.k.StoreProfile(suite.ctx, profile)

	priv1 := secp256k1.GenPrivKey()
	priv2 := secp256k1.GenPrivKey()
	storedLinks := []types.ChainLink{
		types.NewChainLink(
			priv1.PubKey().Address().String(),
			types.NewProof(priv1.PubKey(), "signature", "plain_text"),
			types.NewChainConfig("cosmos", "cosmos"),
			time.Time{},
		),
		types.NewChainLink(
			priv2.PubKey().Address().String(),
			types.NewProof(priv2.PubKey(), "signature", "plain_text"),
			types.NewChainConfig("cosmos", "cosmos"),
			time.Time{},
		),
	}
	for _, link := range storedLinks {
		err := suite.k.StoreChainLink(suite.ctx, profile.GetAddress().String(), link)
		suite.Require().NoError(err)
	}

	usecases := []struct {
		name      string
		req       *types.QueryChainsLinksRequest
		shouldErr bool
		expLen    int
	}{
		{
			name:      "valid request returns no error",
			req:       &types.QueryChainsLinksRequest{},
			shouldErr: false,
			expLen:    2,
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			res, err := suite.k.ChainsLinks(sdk.WrapSDKContext(suite.ctx), uc.req)
			suite.Require().NoError(err)
			suite.Require().NotNil(res)
			suite.Require().Equal(uc.expLen, len(res.ChainsLinks))
		})
	}
}
