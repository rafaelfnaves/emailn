package main

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"emailn/internal/infra/database"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	// middlewares CHI
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var request contract.NewCampaign
		render.DecodeJSON(r.Body, &request)

		id, err := service.Create(request)
		if err != nil {
			if errors.Is(err, internalerrors.ErrInternal) {
				render.Status(r, 500)
			} else {
				render.Status(r, 422)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		render.Status(r, 201)
		render.JSON(w, r, map[string]string{"id": id})
	})

	log.Fatal(http.ListenAndServe(":3000", r))
}
