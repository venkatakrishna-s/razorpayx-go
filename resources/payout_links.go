package resources

import (
	"fmt"

	"github.com/razorpay/razorpayx-go/constants"
	"github.com/razorpay/razorpayx-go/requests"
)

type Payouts_link struct {
	Request *requests.Request
}

//Fetch ...
func (pl *Payouts_link) Fetch(pl_ID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.PAYOUTLINKS_URL, pl_ID)
	return pl.Request.Get(url, data, options)
}

//Create ...
func (pl *Payouts_link) Create(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return pl.Request.Post(constants.PAYOUTLINKS_URL, data, options)
}

//Fetch All
func (pl *Payouts_link) Fetch_All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return pl.Request.Get(constants.PAYOUTLINKS_URL, data, options)
}

//Cancel
func (pl *Payouts_link) Cancel(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/cancel", constants.PAYOUTLINKS_URL, id)
	return pl.Request.Post(url, data, options)
}
