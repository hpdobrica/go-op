package op

import (
	"encoding/json"
	"fmt"
)

func GetItem(itemId string) (ItemWithDetails, error) {

	out, err := executor.RunOp(fmt.Sprintf("op get item %s", itemId), nil)

	var item ItemWithDetails

	if err != nil {
		return item, err
	}

	errJson := json.Unmarshal([]byte(out), &item)

	if errJson != nil {
		return item, errJson
	}

	return item, nil
}
