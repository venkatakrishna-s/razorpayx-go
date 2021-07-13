package resources

import (
	"fmt"

	"github.com/razorpay/razorpayx-go/constants"
	"github.com/razorpay/razorpayx-go/requests"
)

//Contact ...
type Contact struct {
	Request *requests.Request
}

//Fetch ...
func (cont *Contact) Fetch(contactID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.CONTACTS_URL, contactID)
	return cont.Request.Get(url, data, options)
}

//Fetch All...
func (cont *Contact) Fetch_All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return cont.Request.Get(constants.CONTACTS_URL, data, options)
}

//Create ...
func (cont *Contact) Create(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	return cont.Request.Post(constants.CONTACTS_URL, data, options)
}

//Update ...
func (cont *Contact) Update(contactID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.CONTACTS_URL, contactID)

	return cont.Request.Patch(url, data, options)
}

//Activate or Deactivate a contact..

func (cont *Contact) Status(contactID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.CONTACTS_URL, contactID)

	return cont.Request.Patch(url, data, options)
}
