package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpayx-go/constants"
	"github.com/razorpay/razorpayx-go/utils"

	"github.com/stretchr/testify/assert"
)

const Test_pl_id = "test_id"

func TestPlCreate(t *testing.T) {

	teardown, fixture := utils.StartMockServer(constants.PAYOUTLINKS_URL, "fake_payout_link")
	defer teardown()
	contact := map[string]interface{}{
		"name":    "Gaurav Kumar",
		"contact": "912345678",
		"email":   "gaurav.kumar@example.com",
		"type":    "customer",
	}

	req_params := map[string]interface{}{

		"account_number": "fake_account_number",
		"contact":        contact,
		"amount":         1000,
		"currency":       "INR",
		"purpose":        "refund",
		"send_sms":       true,
		"send_email":     true,
	}
	body, err := utils.Client.Payout_link.Create(req_params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}

func TestPlFetch(t *testing.T) {

	url := constants.PAYOUTLINKS_URL + "/" + Test_pl_id
	teardown, fixture := utils.StartMockServer(url, "fake_payout_link")
	defer teardown()
	body, err := utils.Client.Payout_link.Fetch(Test_pl_id, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}

func TestPlFetchAll(t *testing.T) {

	teardown, fixture := utils.StartMockServer(constants.PAYOUTLINKS_URL, "fake_payout_link")
	defer teardown()
	body, err := utils.Client.Payout_link.Fetch_All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)

	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}
