package vehicle

import (
	"net/http"
	"strconv"

	"github.com/Matheoia/vehicle-server/storage"
	"go.uber.org/zap"
)

type DeleteHandler struct {
	store  storage.Store
	logger *zap.Logger
}

func NewDeleteHandler(store storage.Store, logger *zap.Logger) *DeleteHandler {
	return &DeleteHandler{
		store:  store,
		logger: logger.With(zap.String("handler", "delete_vehicles")),
	}
}

func (d *DeleteHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	result, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(rw, "Invalid ID", http.StatusBadRequest)
		return
	}

	contexte := r.Context()

	ok, err := d.store.Vehicle().Delete(contexte, result)

	if err != nil {
		http.Error(rw, "Bonjour", http.StatusInternalServerError)
		return
	}
	if ok {
		rw.WriteHeader(http.StatusNoContent)
		return
	
	}
	rw.WriteHeader(http.StatusNotFound)
}
