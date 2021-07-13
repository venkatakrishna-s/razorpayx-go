package razorpay

import (
	"net/http"
	"time"

	"github.com/razorpay/razorpayx-go/constants"
	"github.com/razorpay/razorpayx-go/requests"
	"github.com/razorpay/razorpayx-go/resources"
)

var Request *requests.Request

type Client struct {
	Contact       *resources.Contact
	Fund_Account  *resources.Fund_Account
	Payouts       *resources.Payouts
	Transaction   *resources.Transaction
	Payout_link   *resources.Payouts_link
	FA_Validation *resources.Validation
}

func NewClient(key string, secret string) *Client {
	auth := requests.Auth{Key: key, Secret: secret}
	httpClient := &http.Client{Timeout: requests.TIMEOUT * time.Second}
	Request = &requests.Request{Auth: auth, HTTPClient: httpClient,

		BaseURL: constants.BASE_URL}

	contact := resources.Contact{Request: Request}
	fa := resources.Fund_Account{Request: Request}
	payouts := resources.Payouts{Request: Request}
	payout_link := resources.Payouts_link{Request: Request}
	transaction := resources.Transaction{Request: Request}
	validation := resources.Validation{Request: Request}

	client := Client{
		Contact:       &contact,
		Fund_Account:  &fa,
		Payouts:       &payouts,
		Payout_link:   &payout_link,
		Transaction:   &transaction,
		FA_Validation: &validation,
	}
	return &client
}

func (client *Client) AddHeaders(headers map[string]string) {
	Request.AddHeaders(headers)
}

func (client *Client) SetTimeout(timeout int16) {
	Request.SetTimeout(timeout)
}
