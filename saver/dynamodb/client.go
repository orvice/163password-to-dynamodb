package dynamodb

type Client struct {
	table string
}

func NewClient() *Client  {
	return new(Client)
}

func(c *Client) SetTableName(table string) {
	c.table = table
}


func(c *Client) Store(email,password string) error{
	return nil
}
