package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpayx-go/constants"
	"github.com/razorpay/razorpayx-go/utils"

	"github.com/stretchr/testify/assert"
)

const Test_fav_id = "test_id"

func TestFavCreate(t *testing.T) {

	teardown, fixture := utils.StartMockServer(constants.FAV_URL, "fake_validation")
	defer teardown()
	fa := map[string]interface{}{
		"id": "fake_fa_id",
	}
	req_params := map[string]interface{}{
		"account_number": "fake_account_number",
		"fund_account":   fa,
		"amount":         100,
		"currency":       "INR",
	}
	body, err := utils.Client.FA_Validation.Create(req_params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}

func TestFavFetch(t *testing.T) {

	url := constants.FAV_URL + "/" + Test_fav_id
	teardown, fixture := utils.StartMockServer(url, "fake_validation")
	defer teardown()
	body, err := utils.Client.FA_Validation.Fetch(Test_fav_id, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}

func TestFavFetchAll(t *testing.T) {

	teardown, fixture := utils.StartMockServer(constants.FAV_URL, "fake_validation")
	defer teardown()
	body, err := utils.Client.FA_Validation.Fetch_All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}
