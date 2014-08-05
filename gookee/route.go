package gookee

type Route struct {
	Url  string
	Act  string
	Ctrl interface{}
}

var routeCollection map[int]Route
var interceptFunc interface{}

func init() {
	routeCollection = make(map[int]Route)
}

func (r Route) Regist() {
	routeCollection[len(routeCollection)] = r
}

func Func(f interface{}) {
	interceptFunc = f
}
