package resources

import (
	"fmt"

	"github.com/razorpay/razorpayx-go/constants"
	"github.com/razorpay/razorpayx-go/requests"
)

type Payouts struct {
	Request *requests.Request
}

//Fetch ...
func (payout *Payouts) Fetch(pout_ID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.PAYOUTS_URL, pout_ID)
	return payout.Request.Get(url, data, options)
}

//Fetch All...
func (payout *Payouts) Fetch_All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return payout.Request.Get(constants.PAYOUTS_URL, data, options)
}

//Create ...
func (payout *Payouts) Create(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return payout.Request.Post(constants.PAYOUTS_URL, data, options)
}

//Cancel a Queued Payout

func (payout *Payouts) Cancel(pout_ID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s/cancel", constants.PAYOUTS_URL, pout_ID)

	return payout.Request.Post(url, data, options)
}
