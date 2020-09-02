package stripeclient_test

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/uwaifo/stripeclient"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no enviroment variable found")
	}

}

func TestClient_Customer(t *testing.T) {
	apiKey := os.Getenv("STRIPE_SECRET_KEY")

	c := stripeclient.Client{Key: apiKey}
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

/* Iterations One

Here we create a standard test function with the testing package
The idea is to attempt to call the stripe api to create a cutomer


1. Create an instance of the Client (struct created in my "strioe" package) c.
2. Create a temporal token
3. Also create a Customer object/instance and pass it our temp token as argument. Handle error
4. If customer is nil, or in otherwords not validly created
4. Check if the customer object we attempted to create is reponding with a json that has an id that is prefixed by "cus_"
*/

/* Iteration Two
1. We add the API Key to instantiating of the Client{Key:"myapikey"}
Note do not expose your secret api. Rather store in enviroment variables
Note: I inroduced the init() function to load my enviromant variable using the godotenv package methods.


*/
