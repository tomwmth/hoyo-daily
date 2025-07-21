package internal

import "encoding/json"

func FromJSON[T any](data []byte) (*T, error) {
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func ToJSON[T any](data *T, pretty bool) ([]byte, error) {
	if pretty {
		return json.MarshalIndent(data, "", "    ")
	} else {
		return json.Marshal(data)
	}
}
