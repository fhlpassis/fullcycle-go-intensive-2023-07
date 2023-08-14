package main

import (
	"net/http"

	"github.com/fhlpassis/fullcycle-go-intensivo-2023-07/internal/entity"
	"github.com/labstack/echo/v4"
)

func main() {
	// chi
	/* r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/order", Order)
	http.ListenAndServe(":3000", r) */

	e := echo.New()
	e.GET("/order", Order)
	e.Logger.Fatal(e.Start(":3000"))
}

/* func Order(w http.ResponseWriter, r *http.Request) {
	order, err := entity.NewOrder("123", 1000, 10)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	order.CalculateFinalPrice()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
} */

func Order(c echo.Context) error {
	order := entity.Order{
		ID:    "1",
		Price: 10,
		Tax:   1,
	}
	err := order.CalculateFinalPrice()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, order)
}
