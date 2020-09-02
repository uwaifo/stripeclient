package stripeclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Customer  . .
type Customer struct {
	ID string `json:"id"`
}

// Client . . .
type Client struct {
	Key string `json:"key"`
}

// Customer . . .
func (c *Client) Customer(token string) (*Customer, error) {
	endpoint := "https://api.stripe.com/v1/customers"
	v := url.Values{}
	v.Set("source", token)

	request, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
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

	var cus Customer
	err = json.Unmarshal(callBody, &cus)
	if err != nil {
		return nil, err
	}
	return &cus, nil

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
