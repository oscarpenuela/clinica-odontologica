package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oscarpenuela/clinica-odontologica.git/internal/dentist"
	"github.com/oscarpenuela/clinica-odontologica.git/pkg/web"
)

type DentistHandler struct {
	DentistService dentist.IService
}

func (h *DentistHandler) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError(err.Error()))
		return 
	}

	dentistFound, err := h.DentistService.GetById(id)
	if err != nil {
		if apiErr, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(apiErr.Status, apiErr)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &dentistFound)
}