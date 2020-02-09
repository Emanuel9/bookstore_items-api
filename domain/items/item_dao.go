package items

import (
	"github.com/Emanuel9/bookstore_items-api/clients/elasticsearch"
	"github.com/Emanuel9/bookstore_utils-go/rest_errors"
	"github.com/pkg/errors"
)

const (
	indexItems = "items"
)
func (i *Item) Save() *rest_errors.RestError {
	result, err := elasticsearch.Client.Index( indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}