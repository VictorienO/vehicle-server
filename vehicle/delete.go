package vehicle

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/VictorienO/vehicle-server/pkg/httputil"
	"github.com/VictorienO/vehicle-server/storage"
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
	id_num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		d.logger.Error(
			"Could not parse id",
			zap.Error(err),
		)
		httputil.ServeError(rw, http.StatusBadRequest, err)
		return
	}
	resp, err := d.store.Vehicle().Delete(r.Context(), id_num)
	if err != nil {
		d.logger.Error(
			"Could not delete vehicle",
			zap.Error(err),
		)
		httputil.ServeError(rw, http.StatusBadRequest, err)
		return
	}
	if !resp {
		fmt.Println("ERREUR 404")
	}
	rw.WriteHeader(http.StatusNoContent)
	//http.Error(rw, "Not Implemented", http.StatusInternalServerError)
}
