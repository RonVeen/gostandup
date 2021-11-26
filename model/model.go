package model

import "fmt"

type Client struct {
	Key  string
	Name string
}

func (c Client) String() string {
	return fmt.Sprintf("Key='%s', Name='%s'", c.Key, c.Name)
}
