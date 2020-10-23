package cli

import (
	"bufio"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/r24zeng/tttgame/x/tttgame/keeper"
	"github.com/r24zeng/tttgame/x/tttgame/types"
	"github.com/cosmos/cosmos-sdk/client/keys"
)

func GetCmdInviteGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "open-game [gameID]",
		Short: "open a new game, waiting for another player",
		Args:  cobra.ExactArgs(1),
		RunE: func(ctx sdk.Contex, k keeper.Keeper, cmd *cobra.Command, args []string) error {
			argsGame := string(args[0])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			fromAddress := cliCtx.GetFromAddress()
			msg := types.NewMsgInviteGame(, argsGame)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			if !k.PlayerExist(ctx, msg.PlayerID) {
				k.CreatePlayer(ctx, msg.PlayerID)
			}
			txBytes := k.Sign(ctx, frommAddress, msg)

			// broadcast to a Tendermint node
			res, err := cliCtx.BroadcastTx(txBytes)
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(res)
		},
	}
}

func GetCmdAcceptGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "accept-game [gameID]",
		Short: "accept invitation, start game",
		Args:  cobra.ExtractArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsGame := string(args[0])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			fromAddress := cliCtx.GetFromAddress()
			msg := types.NewMsgAcceptGame(fromAddress, argsGame)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			if !k.PlayerExist(msg.PlayerID) {
				k.CreatePlayer(msg.PlayerID)
			}
			txBytes := k.Sign(ctx, frommAddress, msg)

			// broadcast to a Tendermint node
			res, err := cliCtx.BroadcastTx(txBytes)
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(res)			
		},
	}
}

func GetCmdPlayGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "play-game [coordinate-X] [coordinate-Y]",
		Short: "play a piece to board",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsGame := string(args[0])
			argsX := int(args[1])
			argsY := int(args[2])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			fromAddress := cliCtx.GetFromAddress()
			msg := types.NewMsgPlayGame(fromAddress, argsGame, argsX, argsY)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			txBytes := k.Sign(ctx, frommAddress, msg)

			// broadcast to a Tendermint node
			res, err := cliCtx.BroadcastTx(txBytes)
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(res)	
		},
	}
}
