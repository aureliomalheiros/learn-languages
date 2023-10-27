package main

import "fmt"

type Client struct {
	accountBalance float32
}

func NewAccount() *Client {
	return &Client {accountBalance: 0}
}

//Method
func (c *Client) simulate(newValue float32) float32 {
	fmt.Printf("The customer has %v in his account\n", c.accountBalance)
	c.accountBalance += newValue

	return c.accountBalance
}

func main(){
	myAccount := Client{accountBalance: 10000.1}
	myAccount.simulate(100.20)
	fmt.Println(myAccount.accountBalance)
}