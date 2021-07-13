package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpayx-go/constants"
	"github.com/razorpay/razorpayx-go/utils"

	"github.com/stretchr/testify/assert"
)

const Test_pout_id = "test_id"

func TestPayoutCreate(t *testing.T) {

	teardown, fixture := utils.StartMockServer(constants.PAYOUTS_URL, "fake_payout")
	defer teardown()

	req_params := map[string]interface{}{

		"account_number":       "fake_account_number",
		"fund_account_id":      "fake_fa_id",
		"amount":               1000000,
		"currency":             "INR",
		"mode":                 "IMPS",
		"purpose":              "refund",
		"queue_if_low_balance": true,
	}
	body, err := utils.Client.Payouts.Create(req_params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}

func TestPayoutFetch(t *testing.T) {

	url := constants.PAYOUTS_URL + "/" + Test_pout_id
	teardown, fixture := utils.StartMockServer(url, "fake_payout")
	defer teardown()
	body, err := utils.Client.Payouts.Fetch(Test_pout_id, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}

func TestPayoutFetchAll(t *testing.T) {

	teardown, fixture := utils.StartMockServer(constants.PAYOUTS_URL, "fake_payout")
	defer teardown()
	body, err := utils.Client.Payouts.Fetch_All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)

	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}
