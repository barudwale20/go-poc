package handlers

import (
	"hello/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (h *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		h.getProducts(rw, req)
		return
	}
	if req.Method == http.MethodPost {
		h.addProduct(rw, req)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *Products) getProducts(rw http.ResponseWriter, req *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	// data, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal json object of products", http.StatusInternalServerError)
	}

	// rw.Write(data)
}

func (h *Products) addProduct(rw http.ResponseWriter, req *http.Request) {
	h.l.Println("Adding a product with POST request!")
	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json!", http.StatusBadRequest)
	}

	h.l.Println("Product: %#v", prod)
	data.AddProduct(prod)
}
