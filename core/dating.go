package core

import (
	"fmt"
    "encoding/json"
    "labix.org/v2/mgo/bson"
)

var AllowedDatings = []string{ "L", "S", "A",
                               "a.q.", "a.q.n.", "p.q.", "p.q.n.",
                               "unknown" }

/**
 * Dating - a structure representing a dating
 */
type Dating struct {
    // ID is created by DB automatically and is only a RO property
    Id bson.ObjectId `bson:"_id"`
    //Id string `bson:"_id"`

    // a dating value is defined (as enum) above 
	Dating string

    // this is description of a dating
	Description string
}

func (d *Dating) String() string {
	return d.Dating
}

func (d *Dating) Display() string {
	s := fmt.Sprintf("%q\n%s\n", d.Dating, d.Description)
	return s
}

// serialize the list of datings into JSON
func DatingsToJson(datings []Dating) (data string, err error) {

    var b []byte
    if b, err = json.Marshal(datings); err != nil {
        return
    }
    data = string(b[:])
    return
}
