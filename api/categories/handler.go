package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	cashTrackHttp "github.com/pawannn/cashtrack/internal/pkg/http"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (cA *CategoriesApi) GetCategories(c *gin.Context) {
	reqID := utils.NewUUID()
	categories, err := cA.categoriesService.GetCategories()
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err)
		return
	}
	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "fetched categories successfully", categories)
}
