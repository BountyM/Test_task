package handler

import (
	"net/http"
	"task1/models"
	"time"

	"github.com/gin-gonic/gin"
)

const N time.Duration = 5 * time.Second

type responseData struct {
	status int
	body   map[string]interface{}
}

func (h *Handler) getResults(c *gin.Context) {

	var input models.Data

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	timeChan := make(chan responseData)

	doneChan := make(chan responseData)

	go func() {
		time.Sleep(N)
		timeChan <- responseData{
			status: http.StatusAccepted,
			body:   gin.H{"message": "Для некоторых из отправлнных чисел подсчет незавершен и находится в процессе."},
		}
	}()

	go func() {
		res, err := h.services.GetResults.GetResults(input.Data)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		doneChan <- responseData{
			status: http.StatusOK,
			body:   gin.H{"data": res},
		}
	}()

	res := responseData{}
	select {
	case res = <-doneChan:
	case res = <-timeChan:
	}
	c.JSON(res.status, res.body)
}
