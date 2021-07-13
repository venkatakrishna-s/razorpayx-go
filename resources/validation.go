package resources

import (
	"fmt"

	"github.com/razorpay/razorpayx-go/constants"
	"github.com/razorpay/razorpayx-go/requests"
)

type Validation struct {
	Request *requests.Request
}

//Fetch ...
func (fav *Validation) Fetch(fav_ID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.FAV_URL, fav_ID)
	return fav.Request.Get(url, data, options)
}

//Fetch All...
func (fav *Validation) Fetch_All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return fav.Request.Get(constants.FAV_URL, data, options)
}

//Create ...
func (fav *Validation) Create(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return fav.Request.Post(constants.FAV_URL, data, options)
}
