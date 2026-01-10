package products 

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1. call the service -> ListProducts
	// 2. Return JSON in an HTTP response
}
