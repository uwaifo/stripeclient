package stripeclient_test

import (
	"flag"
	"log"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/uwaifo/stripeclient"
)

var (
	apiKey string
)

func init() {
	flag.StringVar(&apiKey, "key", "", "test secretkey for stripe API. iF PRESENT, INTEGRATION TEST WILL BE RUN USING THIS KEY")
	if err := godotenv.Load(); err != nil {
		log.Println("no enviroment variable found")
	}

}

func TestClient_Customer(t *testing.T) {
	if apiKey == "" {
		t.Skip("No API key provided")
	}
	//stripeSecretKey := os.Getenv("STRIPE_SECRET_KEY")

	//c := stripeclient.Client{Key: stripeSecretKey}
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

/* Iteration Four
	Long story short we are making the testing more Unit/Modular in nature from being fully intefrated as it currently is now.
	Here we are setting up flags that would come to play when performing our integration test.
	With the understanding that the only time we are going to be needing the secret key is doing the Ingegration test,
	we need not have it availabe while performing other tests(Unit test)

1.	Declare a package variable to hold a string(potential apikey/flag)
2.	In out init() function we set the flag and its required parameters
3. Now in the method we lookup the availability of the secret key which if unavalable would skip the test.
4. If the apikey is available then we would do the integration test(make the call to stripe api)

Note that to accomplish step 4 we have to pass the secret key as an argument to the test commns
IE :go test -v -key=sk_mysecretblablahblah


*/
