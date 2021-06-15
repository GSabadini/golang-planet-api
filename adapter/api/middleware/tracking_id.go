package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type (
	TrackingID string
	//TrackingID struct{}
)

//func NewTrackingID() *TrackingID {
//	return &TrackingID{}
//}

// Execute exports X-Tracking-ID as an http middleware
func (c TrackingID) Execute(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := r.Context()
	id := r.Header.Get("x-tracking-id")
	if id == "" {
		id = uuid.New().String()
	}

	ctx = context.WithValue(ctx, TrackingID("tracking_id"), id)
	r = r.WithContext(ctx)

	w.Header().Set("x-tracking-id", id)
	next.ServeHTTP(w, r)
}
