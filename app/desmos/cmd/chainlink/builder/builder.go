package builder

import (
	"github.com/desmos-labs/desmos/v2/app/desmos/cmd/chainlink/types"
	"github.com/desmos-labs/desmos/v2/x/profiles/client/utils"
)

// ChainLinkJSONBuilder allows to build a ChainLinkJSON instance
type ChainLinkJSONBuilder interface {
	BuildChainLinkJSON(chain types.Chain) (utils.ChainLinkJSON, error)
}

// ChainLinkJSONBuilderProvider allows to provide the provider ChainLinkJSONBuilder implementation based on whether
// it should create the JSON chain link for single or
type ChainLinkJSONBuilderProvider func(isSingleAccount bool) ChainLinkJSONBuilder
