package stripeclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Version . . .
const (
	Version         = "2020-08-27"
	DefaultCurrency = "usd"
)

// Customer  . .
type Customer struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	DefaultSource string `json:"default_source"`
}

// Charge  . . .
type Charge struct {
	ID             string `json:"id"`
	Amount         int    `json:"amount"`
	FailureCode    string `json:"failure_code"`
	FailureMessage string `json:"failure_message"`
	Paid           bool   `json:"paid"`
	Status         string `json:"status"`
}

// Client . . .
type Client struct {
	Key string `json:"key"`
}

// Customer . . .
func (c *Client) Customer(token, email string) (*Customer, error) {
	endpoint := "https://api.stripe.com/v1/customers"
	v := url.Values{}
	v.Set("source", token)
	v.Set("email", email)

	request, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Stripe-Version", Version)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(c.Key, "")

	httpClient := http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	callBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(callBody))
	var cus Customer
	err = json.Unmarshal(callBody, &cus)
	if err != nil {
		return nil, err
	}
	return &cus, nil

}

// Charge . . .
func (c *Client) Charge(customerID string, chargeAmount int) (*Charge, error) {
	endpoint := "https://api.stripe.com/v1/charges"
	v := url.Values{}
	v.Set("customer", customerID)
	v.Set("amount", strconv.Itoa(chargeAmount))

	request, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Stripe-Version", Version)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(c.Key, "")

	httpClient := http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	callBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(callBody))
	var chrg Charge
	err = json.Unmarshal(callBody, &chrg)
	if err != nil {
		return nil, err
	}

	return &chrg, nil

}

/* ONE
Th objective of this Client  is to query the stripe API for a Customers details
There we do the folowing :
1. Create a Customer struct to represent the response we are expecting from the call to the stripe API .
	Note that the customer struct is not EXACTLY representative of the stripe response field. But rather it contaisn only filds that
	are of concern to our use case/business case
2. Create the Client struct to represent our own client instance that is making the API  call .

3. cREATE THE CUTOMER method which

*/

/*TWO
1.	Add stripe api endpont for creating a customer to the Customer method
2.	Create/instatiate a variable to hold url.Values{} (standard library).
 	Using its Set() assign  it a "source"  as its key argument
	and the corespoding value should be the token passed as arg to the Cutomer method
3.	Make a http request object with earlier created value as parameters. Check for errors and return
4.	Set header as app/json for the earlier created request object
5.	Set basic authentication the reciever(Client in this case) Key, and password may be ""
6.	Create a http client
7.	Create a resp,err object based on the httpclient and use the Do() method
	while passing the afor constructed request object and handle errors (just return nil,err)
8.	Now defer res.Body.Close()
9.	Using ioutil, read the response on=bjets Body while passing it to a newly created body object. Handle error
10.	Stringify the body object and print it it statndard output
11.	Declare a Customer varuale
12.	***** . err = Unmarshal the body to the shape of the customer variable created ealier
	(Here we are maping the body to the shape of the Customet struct) . Handle error
13.	Finally return & reference of cutomer variable a,d nil

Note: Read up on step 12

*/

/*Three
Note that stripe endpoints have specific versions usually in a date format
Stripe would automatically assign a version of thier API once they get a request with you secret key.
The only way to change this to send a request to strip
1. The objective of this iteration is to ensure that we are locked into a specific version
	of the stripe API regardless of changes  to the secret keys we may provide.
2. We achieve this by indivating the version in our request header of the Clinet method. This "Stripe-Version" parameter is
	passed a global variable (const) declared outsite the Customer method

*/

/*Iteration Five

Here we are going to try to improve our Customer data struct
1.	Add "Email" field to the Customer and also an email string argument to the Customet method
2.	We also add the email value pair  to our url.Values object
3. 	We add a field to the Cutomer struct to represent the paying customers credit card. This is represented as
	"default_source": "card_1HMri2FODwZN8jDTTKhP1BC9", the json response we get fro the stripe API.

4. Further implimetation in test file
*/

/*Iteration Six
Here we attempt to impliment /create the charge endpoint
1.  We begin by creating the Charge struct with the related fields
2.	We create a Charge method for the Charge struct.This method is prety much similar to the earlier created Customer method with diffrences in the reciever, arguement parameters etc.
3. Take  note to stringify the response body to see possible nested error.

*/
