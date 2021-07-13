package resources

import (
	"fmt"

	"github.com/razorpay/razorpayx-go/constants"
	"github.com/razorpay/razorpayx-go/requests"
)

type Transaction struct {
	Request *requests.Request
}

//Fetch All ...
func (transaction *Transaction) Fetch_All(acc_no string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s?%s", constants.TRANSACTIONS_URL, acc_no)
	return transaction.Request.Get(url, data, options)
}

//Fetch transaction ...
func (transaction *Transaction) Fetch(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.TRANSACTIONS_URL, id)
	return transaction.Request.Get(url, data, options)
}
