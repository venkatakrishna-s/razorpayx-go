package resources

import (
	"fmt"

	"github.com/razorpay/razorpayx-go/constants"
	"github.com/razorpay/razorpayx-go/requests"
)

type Fund_Account struct {
	Request *requests.Request
}

//Fetch ...
func (fa *Fund_Account) Fetch(fa_ID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.FUNDACCOUNTS_URL, fa_ID)
	return fa.Request.Get(url, data, options)
}

//Fetch All ...
func (fa *Fund_Account) Fetch_All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return fa.Request.Get(constants.FUNDACCOUNTS_URL, data, options)
}

//Create ...
func (fa *Fund_Account) Create(data map[string]interface{}, options map[string]string, public bool) (map[string]interface{}, error) {
	//Public API call for PCI DSS Non-Compliant Merchants and account type as card (Payout to Cards)
	if public {
		url := fmt.Sprintf("%s/public", constants.FUNDACCOUNTS_URL)
		return fa.Request.Public(url, data, options)
	} else {
		return fa.Request.Post(constants.FUNDACCOUNTS_URL, data, options)
	}
}

//Activate or Deactivate a fund account ...

func (fa *Fund_Account) Status(fa_ID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.FUNDACCOUNTS_URL, fa_ID)

	return fa.Request.Patch(url, data, options)
}
