package db

import (
	"errors"

	scribble "github.com/nanobox-io/golang-scribble"
	"primitivo.fr/applinh/GoFire/utils"
)

func WriteItem(res string, items map[string]interface{}) (string, error) {
	//toInsert := make(map[string]map[string]string)
	exisingItems := ReadRes(res)
	var uuid string

	if id, exist := items["id"]; exist {
		if _, exist2 := exisingItems[id.(string)]; !exist2 {
			uuid = id.(string)
		} else {
			return "", errors.New("ID_EXIST")
		}
	} else {
		uuid = utils.GenerateUuid()
	}
	db, _ := scribble.New("./dbItems", nil)

	exisingItems[uuid] = items
	db.Write("dbItems", res, exisingItems)
	return uuid, nil
}

func UpdateItem(res string, id string, item map[string]interface{}) error {
	exisingItems := ReadRes(res)
	if _, exist := exisingItems[id]; !exist {
		return errors.New("ID_NO_EXIST")
	}
	exisingItems[id] = item

	db, _ := scribble.New("./dbItems", nil)
	e := db.Write("dbItems", res, exisingItems)

	return e
}

func ReadRes(res string) map[string]map[string]interface{} {
	db, _ := scribble.New("./dbItems", nil)
	value := make(map[string]map[string]interface{})
	db.Read("dbItems", res, &value)
	return value
}
