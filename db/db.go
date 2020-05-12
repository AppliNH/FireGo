package db

import (
	scribble "github.com/nanobox-io/golang-scribble"
	"primitivo.fr/applinh/GoFire/utils"
)

// a fish
type Fish struct{ Name string }

func WriteItem(res string, items map[string]string) string {
	//toInsert := make(map[string]map[string]string)
	exisingItems := ReadRes(res)
	uuid := utils.GenerateUuid()
	
	db, _ := scribble.New("./dbItems", nil)
	
	exisingItems[uuid] = items
	db.Write("dbItems", res, exisingItems)
	return uuid
}

func UpdateItem(res string, id string, item map[string]string) {
	exisingItems := ReadRes(res)
	exisingItems[id] = item
	
	db, _ := scribble.New("./dbItems", nil)
	db.Write("dbItems", res, exisingItems)
}

func ReadRes(res string) map[string]map[string]string {
	db, _ := scribble.New("./dbItems", nil)
	value := make(map[string]map[string]string)
	db.Read("dbItems", res, &value)
	return value
}