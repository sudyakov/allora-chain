package types

import (
	"cosmossdk.io/errors"
	alloraMath "github.com/allora-network/allora-chain/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (msg *CreateNewTopicRequest) Validate(maxStringLen uint64) error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.LossMethod) == 0 || uint64(len(msg.LossMethod)) > maxStringLen {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "loss method invalid")
	}
	if msg.EpochLength <= 0 {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "epoch length must be greater than zero")
	}
	if msg.WorkerSubmissionWindow == 0 {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "worker submission window must be greater than zero")
	}
	if msg.GroundTruthLag < msg.EpochLength {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "ground truth lag cannot be lower than epoch length")
	}
	if msg.WorkerSubmissionWindow > msg.EpochLength {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "worker submission window cannot be higher than epoch length")
	}
	if msg.AlphaRegret.Lte(alloraMath.ZeroDec()) || msg.AlphaRegret.Gt(alloraMath.OneDec()) || validateDec(msg.AlphaRegret) != nil {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "alpha regret must be greater than 0 and less than or equal to 1")
	}
	if msg.PNorm.Lt(alloraMath.MustNewDecFromString("2.5")) || msg.PNorm.Gt(alloraMath.MustNewDecFromString("4.5")) || validateDec(msg.PNorm) != nil {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "p-norm must be between 2.5 and 4.5")
	}
	if msg.Epsilon.Lte(alloraMath.ZeroDec()) || msg.Epsilon.IsNaN() || !msg.Epsilon.IsFinite() {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "epsilon must be greater than 0")
	}
	if uint64(len(msg.Metadata)) > maxStringLen {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "metadata invalid")
	}
	// no validation on AllowNegative because either it is true or false
	// and both are valid values
	//	AllowNegative            bool

	if !isAlloraDecZeroOrLessThanOne(msg.MeritSortitionAlpha) {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "merit sortition alpha must be between 0 and 1 inclusive")
	}
	if !isAlloraDecZeroOrLessThanOne(msg.ActiveInfererQuantile) {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "active inferer quantile must be between 0 and 1 inclusive")
	}
	if !isAlloraDecZeroOrLessThanOne(msg.ActiveForecasterQuantile) {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "active forecaster quantile must be between 0 and 1 inclusive")
	}
	if !isAlloraDecZeroOrLessThanOne(msg.ActiveReputerQuantile) {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "active reputer quantile must be between 0 and 1 inclusive")
	}

	return nil
}
