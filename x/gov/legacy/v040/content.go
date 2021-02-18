// Package v040 is copy-pasted from:
// https://github.com/cosmos/cosmos-sdk/blob/v0.41.1/x/gov/types/content.go
package v040

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)

// Constants pertaining to a Content object
const (
	MaxDescriptionLength int = 5000
	MaxTitleLength       int = 140
)

// Content defines an interface that a proposal must implement. It contains
// information such as the title and description along with the type and routing
// information for the appropriate handler to process the proposal. Content can
// have additional fields, which will handled by a proposal's Handler.
// TODO Try to unify this interface with types/module/simulation
// https://github.com/cosmos/cosmos-sdk/issues/5853
type Content interface {
	GetTitle() string
	GetDescription() string
	ProposalRoute() string
	ProposalType() string
	ValidateBasic() error
	String() string
}

// Handler defines a function that handles a proposal after it has passed the
// governance process.
type Handler func(ctx sdk.Context, content Content) error

// ValidateAbstract validates a proposal's abstract contents returning an error
// if invalid.
func ValidateAbstract(c Content) error {
	title := c.GetTitle()
	if len(strings.TrimSpace(title)) == 0 {
		return sdkerrors.Wrap(types.ErrInvalidProposalContent, "proposal title cannot be blank")
	}
	if len(title) > MaxTitleLength {
		return sdkerrors.Wrapf(types.ErrInvalidProposalContent, "proposal title is longer than max length of %d", MaxTitleLength)
	}

	description := c.GetDescription()
	if len(description) == 0 {
		return sdkerrors.Wrap(types.ErrInvalidProposalContent, "proposal description cannot be blank")
	}
	if len(description) > MaxDescriptionLength {
		return sdkerrors.Wrapf(types.ErrInvalidProposalContent, "proposal description is longer than max length of %d", MaxDescriptionLength)
	}

	return nil
}
