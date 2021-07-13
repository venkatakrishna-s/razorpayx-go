# RazorpayX Go Client 

Golang bindings for interacting with the RazorpayX API


## Usage

You need to setup your key and secret. The instructions can be found here :https://razorpay.com/docs/razorpayx/api/#generate-api-key


```
import (
razorpay "github.com/razorpay/razorpayx-go"
)

client := razorpay.NewClient("<YOUR_API_KEY>", "<YOUR_API_SECRET>")
```
Note: All methods below return a map[string]interface{} and error. All methods accept an extraHeaders param of type map[string]string, allowing you to optinally set extra HTTP headers on the request.

### Contact

Fetch a particular contact
```
body, err := client.Contact.Fetch("<contact_id>" , nil, nil)
```
Fetch all contacts
```
body, err := client.Contact.Fetch_All(nil, nil)
```
Create a Contact
```
data := map[string]interface{}{
  "name":"Gaurav Kumar",
  "email":"gaurav.kumar@example.com",
  "contact":"9123456789",
  "type":"employee",
 }
body, err := client.Contact.Create(data, nil)
```

### Fund Account

Fetch a particular fund account
```
body, err := client.Fund_Account.Fetch("<fa_id>" , nil, nil)
```
Fetch all fund accounts
```
body, err := client.Fund_Account.Fetch_All(nil, nil)
```
Create a fund account
```
bank_acc:=map[string]interface{}{
    "name":"Gaurav Kumar",
    "ifsc":"HDFC0000053",
    "account_number":"765432123456789",
}
data := map[string]interface{}{
  "contact_id":"cont_00000000000001",
  "account_type":"bank_account",
  "bank_account":bank_acc,
}
body, err := client.Fund_Account.Create(data, option, false)
```
Note: Fund account create requires a third boolean parameter, this is to indicate if the API call is public or private

True: Public API call

False: Private API call

The public API call is only for fund account creation for a contact's card if you are a PCI DSS Non-Compliant Merchant.


### Payouts

Fetch a particular payout
 ```
body, err := client.Payout.Fetch("<payout_id>" , nil, nil)
```
Fetch all payouts
```
body, err := client.Payout.Fetch_All(nil, nil)
```
Create a payout
```
data := map[string]interface{}{
  "account_number": "<Razorpay_X_account_number>",
  "fund_account_id": "<fund_account_id>",
  "amount": 1000000,
  "currency": "INR",
  "mode": "IMPS",
  "purpose": "refund",
  "queue_if_low_balance": true,
  "reference_id": "Acme Transaction ID 12345",
}
body, err := client.Payout.Create(data, nil)
```
Note: amount is in paisa


### Fund Account Validation

Fetch a particular validation
 ```
body, err := client.FA_Validation.Fetch("<fav_id>" , nil, nil)
```
Fetch all validations
```
body, err := client.FA_Validation.Fetch_All(nil, nil)
```

Validate a Bank Account
```
fund_account := map[string]interface{}{
		"id": "<fund_account_id>",
	}
	data := map[string]interface{}{
		"account_number": "<Razorpay_X_account_number>",
		"fund_account":   fund_account,
		"amount":         100,
		"currency":       "INR",
	}
	body, err := client.FA_Validation.Create(data, nil)
```
Validate a VPA/UPI id
```
fund_account := map[string]interface{}{
		"id": "<fund_account_id>",
	}
data := map[string]interface{}{
		"account_number": "<Razorpay_X_account_number>",
		"fund_account":   fund_account,
    }
body, err := client.FA_Validation.Create(data, nil)
```
Note: For VPA validation amount is not requried.
