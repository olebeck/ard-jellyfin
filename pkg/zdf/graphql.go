package zdf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var client http.Client

func doGraphql(apiToken string, values map[string]any, data any) error {
	reqBody, _ := json.Marshal(values)
	req, _ := http.NewRequest("POST", "https://api.zdf.de/graphql", bytes.NewReader(reqBody))
	req.Header.Set("api-auth", "Bearer "+apiToken)
	req.Header.Set("content-type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("status code: %d, body: %s", res.StatusCode, string(body))
	}

	body, _ := io.ReadAll(res.Body)
	return json.Unmarshal(body, data)
}
