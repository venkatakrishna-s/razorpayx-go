package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpayx-go/constants"
	"github.com/razorpay/razorpayx-go/utils"

	"github.com/stretchr/testify/assert"
)

const Test_contactid = "test_id"

func TestContactCreate(t *testing.T) {

	teardown, fixture := utils.StartMockServer(constants.CONTACTS_URL, "fake_contact")
	defer teardown()
	req_params := map[string]interface{}{
		"name":         "testname",
		"contact":      "9123456789",
		"email":        "test@example.com",
		"type":         "test",
		"reference_id": "testreference",
	}
	body, err := utils.Client.Contact.Create(req_params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}

func TestContactFetch(t *testing.T) {

	url := constants.CONTACTS_URL + "/" + Test_contactid
	teardown, fixture := utils.StartMockServer(url, "fake_contact")
	defer teardown()
	body, err := utils.Client.Contact.Fetch(Test_contactid, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}

func TestContactFetchAll(t *testing.T) {

	url := constants.CONTACTS_URL
	teardown, fixture := utils.StartMockServer(url, "fake_contact")
	defer teardown()
	body, err := utils.Client.Contact.Fetch_All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}

func TestContactUpdate(t *testing.T) {
	url := constants.CONTACTS_URL + "/" + Test_contactid
	teardown, fixture := utils.StartMockServer(url, "fake_contact")
	defer teardown()
	req_params := map[string]interface{}{
		"name": "testname",
	}
	body, err := utils.Client.Contact.Update(Test_contactid, req_params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
