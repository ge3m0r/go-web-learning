package v1

import (
    "golearning/app/models/link"
    "golearning/pkg/response"

    "github.com/gin-gonic/gin"
)

type LinksController struct {
    BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
    response.Data(c, link.AllCached())
}

