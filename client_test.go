package stripeclient_test

import (
	"strings"
	"testing"

	"github.com/uwaifo/stripeclient"
)

func TestClient_Customer(t *testing.T) {

	c := stripeclient.Client{}
	tok := "tok_amex"
	cus, err := c.Customer(tok)
	if err != nil {
		t.Errorf("Customer() err = %v; Wants %v", err, nil)
	}

	if cus == nil {
		t.Fatal("Customer() = nil; wants non-nil value")
	}

	if !strings.HasPrefix(cus.ID, "cus_") {
		t.Errorf("Customer() ID = %v; Wants prefix %q", cus.ID, "cus_")
	}

}

/*
Here we create a standard test function with the testing package
The idea is to attempt to call the stripe api to create a cutomer

curl https://api.stripe.com/v1/customers \
  -u sk_test_pflyokeRxblxT6MkGjzIEybT005avqxgI6: \
  -d description="My First Test Customer (created for API docs)"



1. Create an instance of the Client (struct created in my "strioe" package) c.
2. Create a temporal token
3. Also create a Customer object/instance and pass it our temp token as argument. Handle error
4. If customer is nil, or in otherwords not validly created
4. Check if the customer object we attempted to create is reponding with a json that has an id that is prefixed by "cus_"
*/
