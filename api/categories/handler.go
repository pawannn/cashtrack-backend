package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	cashTrackHttp "github.com/pawannn/cashtrack/internal/pkg/http"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (cA *CategoriesApi) GetCategories(c *gin.Context) {
	reqID := utils.NewUUID()

	cA.categoriesLogger.Info(reqID, "Fetching all categories")

	categories, err := cA.categoriesService.GetCategories()
	if err != utils.NoErr {
		cA.categoriesLogger.Error(reqID, "Failed to fetch categories", err.Error)
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err)
		return
	}

	cA.categoriesLogger.Info(reqID, "Fetched categories successfully", len(categories))
	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "Fetched categories successfully", categories)
}
