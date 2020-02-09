package controllers

import (
	"encoding/json"
	"github.com/Emanuel9/bookstore_items-api/domain/items"
	"github.com/Emanuel9/bookstore_items-api/logger"
	"github.com/Emanuel9/bookstore_items-api/services"
	"github.com/Emanuel9/bookstore_items-api/utils/http_utils"
	"github.com/Emanuel9/bookstore_oauth-go/oauth"
	"github.com/Emanuel9/bookstore_utils-go/rest_errors"
	"io/ioutil"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter,*http.Request)
	Get(http.ResponseWriter,*http.Request)
}

type itemsController struct{}

func (c *itemsController ) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//http_utils.RespondError(w, *err) -- outh does not use shared library yet
		logger.Info("Intru aici")
		return
	}

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("invalid request body")
		http_utils.RespondError(w, *respErr)
		return
	}

	var itemRequest items.Item

	requestBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, *respErr)
		return
	}

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, *respErr)
		return
	}

	itemRequest.Seller = sellerId
	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, *createErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}