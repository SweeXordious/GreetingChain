package helloworld

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sweexordious/helloworld/x/helloworld/types"
)

// NewHandler creates an sdk.Handler for all the helloworld type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgSet:
			return handleMsgSet(ctx, k, msg)
		case MsgGet:
			return handleMsgGet(ctx, k, msg)
		case MsgPropose:
			return handleMsgPropose(ctx, k, msg)
		case MsgSell:
			return handleMsgSell(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handle<Action> does x
func handleMsgSet(ctx sdk.Context, k Keeper, msg MsgSet) (*sdk.Result, error) {
	err := k.SetMsg(ctx, types.Hello{
		Owner: msg.Sender,
		Msg:   msg.Hello,
		Price: msg.Price,
	})

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
			sdk.NewAttribute("Msg: ", msg.Hello),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// handle<Action> does x
func handleMsgGet(ctx sdk.Context, k Keeper, msg MsgGet) (*sdk.Result, error) {
	message, err := k.GetMsg(ctx, msg.Hello)
	if err != nil {
		fmt.Printf("could not get hello in handleMsgGet\n%s\n", err.Error())
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.ValidatorAddr.String()),
			sdk.NewAttribute("Msg: ", message.Msg),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgPropose(ctx sdk.Context, k Keeper, msg MsgPropose) (*sdk.Result, error) {
	err := k.ProposeMsg(ctx, types.Hello{
		Owner: msg.Sender,
		Msg:   msg.Hello,
		Price: msg.Price,
	})
	if err != nil {
		fmt.Printf("could not buy hello in handleMsgGet\n%s\n", err.Error())
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
			sdk.NewAttribute("Msg proposal: ", msg.Hello),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgSell(ctx sdk.Context, k Keeper, msg MsgSell) (*sdk.Result, error) {
	err := k.SellMsg(ctx, types.Hello{
		Owner: msg.Sender,
		Msg:   msg.Hello,
	})
	if err != nil {
		fmt.Printf("could not sell msg \n%s\n", err.Error())
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
			sdk.NewAttribute("Msg sell: ", msg.Hello),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
