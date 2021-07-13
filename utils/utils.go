package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	razorpay "github.com/razorpay/razorpayx-go"

	"github.com/stretchr/testify/assert"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	Client *razorpay.Client
)

const (
	TestAPIKey    = "fake_key"
	TestAPISecret = "fake_secret"
)

func testSetup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	Client = razorpay.NewClient(TestAPIKey, TestAPISecret)
	razorpay.Request.BaseURL = server.URL

	return func() {
		server.Close()
	}
}

func getFixture(path string) string {
	filepath := path + ".json"
	b, err := ioutil.ReadFile("../test_data/" + filepath)
	if err != nil {
		panic(err)
	}

	return strings.Replace(strings.Replace(string(b), "\n", "", -1), " ", "", -1)
}

func jsonCompare(b1, b2 []byte) (bool, error) {
	var o1, o2 interface{}
	err := json.Unmarshal(b1, &o1)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(b2, &o2)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(o1, o2), nil
}

//TestResponse ...
func TestResponse(jsonBodyByteArray []byte, fixtureByteArray []byte, t *testing.T) {
	status, err := jsonCompare(fixtureByteArray, jsonBodyByteArray)
	if err == nil {
		assert.Equal(t, status, true)
	} else {
		panic(err.Error())
	}
}

//StartMockServer ...
func StartMockServer(url string, fixtureName string) (func(), string) {
	teardown := testSetup()
	fixture := getFixture(fixtureName)
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, fixture)
	})
	return teardown, fixture
}
