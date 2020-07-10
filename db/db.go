package db

import (
	scribble "github.com/nanobox-io/golang-scribble"
	"primitivo.fr/applinh/GoFire/utils"
)

func WriteItem(res string, items map[string]interface{}) string {
	//toInsert := make(map[string]map[string]string)
	exisingItems := ReadRes(res)
	var uuid string

	if id, exist := items["id"]; exist {
		if _, exist2 := exisingItems[id.(string)]; !exist2 {
			uuid = id.(string)
		}
	} else {
		uuid = utils.GenerateUuid()
	}
	db, _ := scribble.New("./dbItems", nil)

	exisingItems[uuid] = items
	db.Write("dbItems", res, exisingItems)
	return uuid
}

func UpdateItem(res string, id string, item map[string]interface{}) {
	exisingItems := ReadRes(res)
	exisingItems[id] = item

	db, _ := scribble.New("./dbItems", nil)
	db.Write("dbItems", res, exisingItems)
}

func ReadRes(res string) map[string]map[string]interface{} {
	db, _ := scribble.New("./dbItems", nil)
	value := make(map[string]map[string]interface{})
	db.Read("dbItems", res, &value)
	return value
}
