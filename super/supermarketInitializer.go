package super

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"errors"
)

type App struct {
	Router *mux.Router
	Stock  Stock
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (app *App) initializeRoutes() {
	app.Router = mux.NewRouter().StrictSlash(true)
	app.Router.HandleFunc("/stock/{name:[a-z]+}", app.buy).Methods("POST")
	app.Router.HandleFunc("/stock/products", app.getProducts)

}

func (a *App) Run(addr string) {
	a.Stock = initSupermarket()
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func initSupermarket() Stock {
	return Stock{Contents: []ProductSelection{
		{Prod: Product{Name: "eau", Price: 1}, Amt: 50},
		{Prod: Product{Name: "gatorade", Price: 2}, Amt: 10},
		{Prod: Product{Name: "stickers", Price: 10, Discount: 3}, Amt: 50},
		{Prod: Product{Name: "chocolat", Price: 15, Discount: 1}, Amt: 10},
		{Prod: Product{Name: "crepes", Price: 1}, Amt: 100},
		{Prod: Product{Name: "gelatine", Price: 3}, Amt: 30},
		{Prod: Product{Name: "piles", Price: 10}, Amt: 10},
		{Prod: Product{Name: "nivea", Price: 20, Discount: 2}, Amt: 50}}}
}

func (app *App) getProducts(writer http.ResponseWriter, request *http.Request) {
	respondWithJSON(writer, http.StatusOK, app.Stock.Contents)
}

func (app *App) buy(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]

	product, err := app.findProduct(name)
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, "Error finding a product")
		return
	}
	var p ProductSelection
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(writer, http.StatusBadRequest, "Invalid request amount")
		return
	}
	defer request.Body.Close()

	app.Stock.SellProduct(ProductSelection{Prod: product, Amt: p.Amt})
	respondWithJSON(writer, http.StatusCreated, p)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (app *App) findProduct(name string) (Product, error) {
	for _, content := range app.Stock.Contents {
		fmt.Println(content.Prod.Name)
		if content.Prod.Name == name {
			return content.Prod, nil
		}
	}
	return Product{}, errors.New("not found")
}
