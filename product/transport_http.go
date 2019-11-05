package product

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getProductByIdHandler := kithttp.NewServer(makeGetProductByIdEndPoint(s),
		getProductByIdRequestDecoder,
		kithttp.EncodeJSONResponse)

	r.Method(http.MethodGet, "/{id}", getProductByIdHandler)

	return r
}
func getProductByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIDRequest{
		ProductID: productId,
	}, nil
}
