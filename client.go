package stripeclient

// Customer  . .
type Customer struct {
	ID string `json:"id"`
}

// Client . . .
type Client struct{}

// Customer . . .
func (c *Client) Customer(token string) (*Customer, error) {
	return nil, nil

}

/*
Th objective of this Client  is to query the stripe API for a Customers details
There we do the folowing :
1. Create a Customer struct to represent the response we are expecting from the call to the stripe API .
	Note that the customer struct is not EXACTLY representative of the stripe response field. But rather it contaisn only filds that
	are of concern to our use case/business case
2. Create the Client struct to represent our own client instance that is making the API  call .

3. cREATE THE CUTOMER method which

*/
