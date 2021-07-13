package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpayx-go/constants"
	"github.com/razorpay/razorpayx-go/utils"

	"github.com/stretchr/testify/assert"
)

const Test_txn_id = "test_id"
const Test_acc_no = "test_no"

func TestTxnFetch(t *testing.T) {

	url := constants.TRANSACTIONS_URL + "/" + Test_txn_id
	teardown, fixture := utils.StartMockServer(url, "fake_transaction")
	defer teardown()
	body, err := utils.Client.Transaction.Fetch(Test_txn_id, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}

func TestTxnFetchAll(t *testing.T) {

	teardown, fixture := utils.StartMockServer(constants.TRANSACTIONS_URL, "fake_transaction")
	defer teardown()
	body, err := utils.Client.Transaction.Fetch_All(Test_acc_no, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)

}
