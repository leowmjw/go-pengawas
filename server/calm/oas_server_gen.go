// Code generated by ogen, DO NOT EDIT.

package calm

import (
	"context"

	"go.opentelemetry.io/otel/metric/instrument/syncint64"

	"github.com/ogen-go/ogen/otelogen"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// EnrollCard implements enrollCard operation.
	//
	// Enrolls a card to card lifecycle management system for future account update notifications.
	// CALM will perform both card availability and eligibility checks to check that this specific card
	// is eligible for management.
	// Only Mastercard cards are supported right now.
	//
	// POST /cards
	EnrollCard(ctx context.Context, req OptCard) (EnrollCardRes, error)
	// GetCard implements getCard operation.
	//
	// Returns information of enrolled card with the most up-to-date card information available from
	// issuing banks and all CALM capabilities.
	//
	// GET /cards/{cardId}
	GetCard(ctx context.Context, params GetCardParams) (GetCardRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	cfg config

	requests syncint64.Counter
	errors   syncint64.Counter
	duration syncint64.Histogram
}

func NewServer(h Handler, opts ...Option) (*Server, error) {
	s := &Server{
		h:   h,
		cfg: newConfig(opts...),
	}
	var err error
	if s.requests, err = s.cfg.Meter.SyncInt64().Counter(otelogen.ServerRequestCount); err != nil {
		return nil, err
	}
	if s.errors, err = s.cfg.Meter.SyncInt64().Counter(otelogen.ServerErrorsCount); err != nil {
		return nil, err
	}
	if s.duration, err = s.cfg.Meter.SyncInt64().Histogram(otelogen.ServerDuration); err != nil {
		return nil, err
	}
	return s, nil
}
