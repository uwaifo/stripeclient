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
	email := "glogussee@gmail.com"
	cus, err := c.Customer(tok, email)
	if err != nil {
		t.Errorf("Customer() err = %v; Wants %v", err, nil)
	}

	if cus == nil {
		t.Fatal("Customer() = nil; wants non-nil value")
	}

	if !strings.HasPrefix(cus.ID, "cus_") {
		t.Errorf("Customer() ID = %v; Wants prefix %q", cus.ID, "cus_")
	}

	// check card format validity (Not actually validataed card but just the format)
	if !strings.HasPrefix(cus.DefaultSource, "card_") {
		t.Errorf("Customer() DefaultStore= %v; Wants prefix %q", cus.DefaultSource, "card_")
	}

	// Check email
	if cus.Email != email {
		t.Errorf("Customer() Email= %v; Wants %q", cus.Email, email)
	}

}

func TestClient_Charge(t *testing.T) {
	//API Key validation section
	if apiKey == "" {
		t.Skip("No API key provided")
	}

	c := stripeclient.Client{Key: apiKey}

	//Create a test customer to be charged
	tok := "tok_amex"
	email := "glogussee@gmail.com"
	cus, err := c.Customer(tok, email)
	if err != nil {
		t.Errorf("Customer() err = %v; Wants nil", err)
	}
	//

	//Test Parameters
	_ = cus
	amount := 8800
	charge, err := c.Charge("cus_ID", amount)

	if err != nil {
		t.Errorf("Customer() err = %v; Wants %v", err, nil)
	}

	if charge == nil {
		t.Fatal("Charge() = nil; wants non-nil value")
	}

	// Check amount
	if charge.Amount != amount {
		t.Errorf("Charge() Amount= %d; Wants %d", charge.Amount, amount)
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

/*Iteration Five
Here we are going to try to improve our Customer data struct
1.	Add "Email" field to the Customer and also an email string argument to the Customet method
2.	We also add the email value pair  to our url.Values object
3. 	We add a field to the Cutomer struct to represent the paying customers credit card. This is represented as
	"default_source": "card_1HMri2FODwZN8jDTTKhP1BC9", the json response we get fro the stripe API.
4.	Further implimetation in test file
5.	Now that we habe introduced the Email and DefautSource to the Cutomer struct and the same is required
	as parameters of the Customer method we have to expect then in the client test also.
6. 	Add validation check (if block) for the email and datasource fileds

*/

/*Iteration Six
Create the TestClient_Charge function . Similar to TestClient_Customer.
1. Note the usage of "_ = cus" to prevent go from complaining about not using the Customer cus object while we test witch some dummy data.


*/
