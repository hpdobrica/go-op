package op

import (
	"fmt"
)

func GetItem(itemId string) (string, error) {

	out, err := executor.RunOp(fmt.Sprintf("op get item %s", itemId), nil)

	if err != nil {
		return "", err
	}

	// var items []Item
	// errJson := json.Unmarshal([]byte(out), &items)

	// if errJson != nil {
	// 	return "", errJson
	// }

	return out, nil
}
