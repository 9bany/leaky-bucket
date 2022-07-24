package rule

type Rule struct {
	Amout int64
}

/*
* These rules map can be fetched form a database rather than
* herdcoding at application layer and also can be periodically
* update in a background job in case some dynamic changed are needed.
 */
var rulesMap = map[string]*Rule{
	"gen-user": {Amout: (1 << 20)},
}

func GetBucket(userType string) *Rule {
	return rulesMap[userType]
}
