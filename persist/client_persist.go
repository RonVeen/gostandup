package persist

import (
	"bufio"
	"errors"
	"fmt"
	"gostandup/model"
	"os"
	"strings"
)

func AppendClient(c model.Client) (bool, error) {
	f, err := os.OpenFile("client.data",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false, err
	}
	defer f.Close()

	s := fmt.Sprintf(c.Key + "," + c.Name + "\n")
	if _, err := f.WriteString(s); err != nil {
		return false, err
	}

	return true, nil
}

func ReadClient() ([]model.Client, error) {
	f, err := os.Open("client.data")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	results := make([]model.Client, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		c, err := assembleClient(scanner.Text())
		if err == nil {
			results = append(results, *c)
		}
	}

	return results, nil
}

func assembleClient(input string) (*model.Client, error) {
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return nil, errors.New("Invalid number of elements: " + input)
	}
	return &model.Client{
		Key:  parts[0],
		Name: parts[1],
	}, nil
}
