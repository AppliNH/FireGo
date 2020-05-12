package db

import (
	"encoding/json"
	"fmt"

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

func DbProcess() {
	db, _ := scribble.New("./dbItems", nil)

	// add some fish to the database
	myFishies := []Fish{{Name: "onefish"}, {Name: "twofish"}, {Name: "redfish"}, {Name: "bluefish"}}
	db.Write("fish", "fish", myFishies)
	// Read one fish from the database
	onefish := Fish{}
	db.Read("fish", "onefish", &onefish)

	fmt.Printf("It's a fish! %#v\n", onefish)

	// Read more fish from the database
	morefish, _ := db.ReadAll("fish")

	// iterate over morefish creating a new fish for each record
	fishies := []Fish{}
	for _, fish := range morefish {
		f := Fish{}
		json.Unmarshal([]byte(fish), &f)
		fishies = append(fishies, f)
	}

	fmt.Printf("It's a lot of fish! %#v\n", fishies)

	// Delete onefish from the database
	//db.Delete("fish", "onefish")

	// Delete all fish from the database
	//db.Delete("fish", "")
}
