package app

import "github.com/gin-gonic/gin"

type ResponseJson struct {
	Status int
	Data   interface{}
}

func RespondJSON(w *gin.Context, status int, payload interface{}) {
	var res ResponseJson
	res.Status = status
	res.Data = payload
	w.JSON(200, res)
}
