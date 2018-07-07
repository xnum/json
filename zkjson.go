package json

import (
	"encoding/json"
)

// Rawer interface is to get internal data.
type Rawer interface {
	Raw() interface{}
}

// ObjPayload stores members.
type ObjPayload struct {
	Attr map[string]interface{}
}

// Members stores key-value pair temporarily.
type Member struct {
	Key string
	Val interface{}
}

// Object returns a new ObjPayload.
func Object(arr ...Member) *ObjPayload {
	obj := &ObjPayload{Attr: make(map[string]interface{})}
	obj.Append(arr...)
	return obj
}

// Append adds member to ObjPayload.
func (p *ObjPayload) Append(arr ...Member) *ObjPayload {
	for _, e := range arr {
		p.Attr[e.Key] = e.Val
	}
	return p
}

// Marshal returns its json string.
func (p *ObjPayload) Marshal() ([]byte, error) {
	return json.Marshal(p.Attr)
}

// Raw returns its raw data.
func (p *ObjPayload) Raw() interface{} {
	return p.Attr
}

// Attr accepts key-value pair and return Member object.
// it converts `key, val` to `key: val`
func Attr(key string, val interface{}) Member {
	if raw, ok := val.(Rawer); ok {
		return Member{Key: key, Val: raw.Raw()}
	}
	if m, ok := val.(Member); ok {
		return Member{Key: key, Val: Object(Attr(m.Key, m.Val)).Raw()}
	}
	return Member{Key: key, Val: val}
}

// ArrPayload stores values.
type ArrPayload struct {
	eles []interface{}
}

// Marshal returns its json string.
func (a *ArrPayload) Marshal() ([]byte, error) {
	return json.Marshal(a.eles)
}

// Raw returns its raw data.
func (a *ArrPayload) Raw() interface{} {
	return a.eles
}

// Array stores values and returns a new ArrPayload. it converts to `[]`
func Array(arr ...interface{}) *ArrPayload {
	ret := &ArrPayload{}
	if len(arr) == 0 {
		ret.eles = make([]interface{}, 0)
	} else {
		ret.Append(arr...)
	}
	return ret
}

// Append appends anything as its values.
func (a *ArrPayload) Append(arr ...interface{}) *ArrPayload {
	for _, v := range arr {
		if raw, ok := v.(Rawer); ok {
			a.eles = append(a.eles, raw.Raw())
		} else {
			a.eles = append(a.eles, v)
		}
	}
	return a
}
