package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"workshop/internal/domain/sighting"
	"workshop/web/views/pages"
)

type SightingsLister interface {
	List(ctx context.Context, filter sighting.ListFilter) ([]sighting.Sighting, error)
}

func Home(logger *slog.Logger, lister SightingsLister) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement
		// 1. Call the sightings service to List.
		// 2. Set content type to text/html.
		// 3. Render home page
		ctx := r.Context()

		list, err := lister.List(ctx, sighting.ListFilter{
			Limit: 10,
		})
		if err != nil {
			return
		}

		err = pages.Home(list).Render(ctx, w)
		if err != nil {
			return
		}

	}
}
