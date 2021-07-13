package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpayx-go/constants"
	"github.com/razorpay/razorpayx-go/utils"

	"github.com/stretchr/testify/assert"
)

const Test_fa_id = "test_id"

func TestFaCreate(t *testing.T) {

	teardown, fixture := utils.StartMockServer(constants.FUNDACCOUNTS_URL, "fake_fund_account")
	defer teardown()
	bank_acc := map[string]interface{}{

		"name":           "Gaurav Kumar",
		"ifsc":           "HDFC0000053",
		"account_number": "765432123456789",
	}
	req_params := map[string]interface{}{

		"contact_id":   "fake_contactid",
		"account_type": "bank_account",
		"bank_account": bank_acc,
	}
	body, err := utils.Client.Fund_Account.Create(req_params, nil, false)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}

func TestFaFetch(t *testing.T) {

	url := constants.FUNDACCOUNTS_URL + "/" + Test_fa_id
	teardown, fixture := utils.StartMockServer(url, "fake_fund_account")
	defer teardown()
	body, err := utils.Client.Fund_Account.Fetch(Test_contactid, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)

	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}

func TestFaFetchAll(t *testing.T) {

	teardown, fixture := utils.StartMockServer(constants.FUNDACCOUNTS_URL, "fake_fund_account")
	defer teardown()
	body, err := utils.Client.Fund_Account.Fetch_All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)

	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}
