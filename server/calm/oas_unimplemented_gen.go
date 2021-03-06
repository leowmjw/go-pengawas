// Code generated by ogen, DO NOT EDIT.

package calm

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

var _ Handler = UnimplementedHandler{}

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

// EnrollCard implements enrollCard operation.
//
// Enrolls a card to card lifecycle management system for future account update notifications.
// CALM will perform both card availability and eligibility checks to check that this specific card
// is eligible for management.
// Only Mastercard cards are supported right now.
//
// POST /cards
func (UnimplementedHandler) EnrollCard(ctx context.Context, req OptCard) (r EnrollCardRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetCard implements getCard operation.
//
// Returns information of enrolled card with the most up-to-date card information available from
// issuing banks and all CALM capabilities.
//
// GET /cards/{cardId}
func (UnimplementedHandler) GetCard(ctx context.Context, params GetCardParams) (r GetCardRes, _ error) {
	return r, ht.ErrNotImplemented
}
