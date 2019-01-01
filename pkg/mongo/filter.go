package mongo

import (
	"net/url"

	"github.com/mongodb/mongo-go-driver/bson"
)

// FiltersFromURLValues does not handle anything else except
// equals value, nor and/or filters for now.
func FiltersFromURLValues(urlValues url.Values) bson.M {
	doc := bson.M{}
	for k, values := range urlValues {
		if len(values) == 0 {
			continue
		}
		doc[k] = values[0]
	}
	return doc
}
