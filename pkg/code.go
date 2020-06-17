package pkg

var Codes = map[int]string{
	500:   "system error",
	10000: "token error",
	10001: "data layer error",
	10002: "params error",
	10003: "no authorization",
	20000: "no app authorization",
	20001: "no business authorization",
	20002: "business not found",
	20003: "store not found",
}
