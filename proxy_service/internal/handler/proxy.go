package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/roxyash/kmf_testtask/proxy_service/internal/request"
	"github.com/roxyash/kmf_testtask/proxy_service/internal/response"
	"io/ioutil"
	"net/http"
	"strings"
)

// @Summary	Proxy
// @Tags		Proxy
// @Accept		json
// @Produce	json
// @Param		user	body		request.ProxyRequest	true	"User login details"
// @Success	200		{object}	response.ProxyResponse
// @Failure	400		{string}	json				"Error"
// @Router		/proxy [post]
func (h *Handler) Proxy(ctx *gin.Context) {
	var requestBody request.ProxyRequest

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest(requestBody.Method, requestBody.URL, nil)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	for key, value := range requestBody.Headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": "Bad gateway"})
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": "Bad gateway"})
		return
	}

	headers := make(map[string]string)
	for key, values := range resp.Header {
		headers[key] = strings.Join(values, ",")
	}

	//Generate a unique ID for the response and store it in the response map
	proxyResponse := h.service.Proxy.SetProxyResponseData(response.ProxyResponse{
		Status:  resp.StatusCode,
		Headers: headers,
		Length:  int64(len(respBody)),
	})

	// Return the response ID to the client
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, proxyResponse)
}
