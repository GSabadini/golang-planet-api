package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type TrackingID struct{}

func NewTrackingID() *TrackingID {
	return &TrackingID{}
}

func (c TrackingID) Execute(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := r.Context()
	id := r.Header.Get("X-Tracking-Id")
	if id == "" {
		id = uuid.New().String()
	}

	ctx = context.WithValue(ctx, "tracking_id", id)
	r = r.WithContext(ctx)

	w.Header().Set("X-Tracking-Id", id)
	next.ServeHTTP(w, r)
}
