package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

// GetCmdCreateRelationship returns the command allowing to create a relationship
func GetCmdCreateRelationship() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-relationship [receiver] [subspace-id]",
		Short: "Create a relationship with the given receiver address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRelationship(clientCtx.FromAddress.String(), args[0], args[1])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdDeleteRelationship returns the command allowing to delete a relationships
func GetCmdDeleteRelationship() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-relationship [receiver] [subspace-id]",
		Short: "Delete the relationship with the given user",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteRelationship(clientCtx.FromAddress.String(), args[0], args[1])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdBlockUser returns the command allowing to block a user
func GetCmdBlockUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block [address] [subspace] [[reason]]",
		Short: "Block the user with the given address, optionally specifying the reason for the block",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			reason := ""
			if len(args) == 3 {
				reason = args[2]
			}

			msg := types.NewMsgBlockUser(clientCtx.FromAddress.String(), args[0], reason, args[1])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdUnblockUser returns the command allowing to unblock a user
func GetCmdUnblockUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unblock [address] [subspace]",
		Short: "Unblock the user with the given address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUnblockUser(clientCtx.FromAddress.String(), args[0], args[1])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// --------------------------------------------------------------------------------------------------------------------

// GetCmdQueryRelationships returns the command allowing to query the relationships with optional user and subspace
func GetCmdQueryRelationships() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relationships [[address]] [[subspace-id]]",
		Short: "Retrieve all the relationships with optional address and subspace",
		Args:  cobra.RangeArgs(0, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			var user string
			if len(args) >= 1 {
				user = args[0]
			}

			var subspace string
			if len(args) == 2 {
				subspace = args[1]
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Relationships(
				context.Background(),
				&types.QueryRelationshipsRequest{User: user, SubspaceId: subspace, Pagination: pageReq},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "user relationships")

	return cmd
}

// GetCmdQueryBlocks returns the command allowing to query all the blocks with optional user and subspace
func GetCmdQueryBlocks() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blocks [[address]] [[subspace-id]] ",
		Short: "Retrieve the list of all the blocked users with optional address and subspace",
		Args:  cobra.RangeArgs(0, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			var user string
			if len(args) >= 1 {
				user = args[0]
			}

			var subspace string
			if len(args) == 2 {
				subspace = args[1]
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Blocks(
				context.Background(),
				&types.QueryBlocksRequest{User: user, SubspaceId: subspace, Pagination: pageReq})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "user blocks")

	return cmd
}
