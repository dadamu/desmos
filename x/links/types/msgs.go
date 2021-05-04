package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgCreateIBCAccountConnection(
	port string,
	channelID string,
	timeoutTimestamp uint64,
	sourceAddress string,
	sourcePubKey string,
	destinationAddress string,
	sourceSignature string,
	destinationSignature string,
) *MsgCreateIBCAccountConnection {
	return &MsgCreateIBCAccountConnection{
		Port:                 port,
		ChannelId:            channelID,
		TimeoutTimestamp:     timeoutTimestamp,
		SourceAddress:        sourceAddress,
		SourcePubKey:         sourcePubKey,
		DestinationAddress:   destinationAddress,
		SourceSignature:      sourceSignature,
		DestinationSignature: destinationSignature,
	}
}

// Route should return the name of the module
func (msg MsgCreateIBCAccountConnection) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateIBCAccountConnection) Type() string { return ActionIBCAccountConnection }

// ValidateBasic runs stateless checks on the message
func (msg *MsgCreateIBCAccountConnection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.SourceAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid source address (%s)", err)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateIBCAccountConnection) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgCreateIBCAccountConnection) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.SourceAddress)
	return []sdk.AccAddress{sender}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgCreateIBCAccountConnection) MarshalJSON() ([]byte, error) {
	type temp MsgCreateIBCAccountConnection
	return json.Marshal(temp(msg))
}

// ___________________________________________________________________________________________________________________

func NewMsgCreateIBCAccountLink(
	port string,
	channelID string,
	timeoutTimestamp uint64,
	sourceAddress string,
	sourcePubKey string,
	signature string,
) *MsgCreateIBCAccountLink {
	return &MsgCreateIBCAccountLink{
		Port:             port,
		ChannelId:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		SourceAddress:    sourceAddress,
		SourcePubKey:     sourcePubKey,
		Signature:        signature,
	}
}

// Route should return the name of the module
func (msg MsgCreateIBCAccountLink) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateIBCAccountLink) Type() string { return ActionIBCAccountLink }

// ValidateBasic runs stateless checks on the message
func (msg *MsgCreateIBCAccountLink) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.SourceAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid source address (%s)", err)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateIBCAccountLink) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgCreateIBCAccountLink) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.SourceAddress)
	return []sdk.AccAddress{sender}
}