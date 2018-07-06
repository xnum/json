package json

import (
	"fmt"
	"strings"
)

// ObjPayload stores members.
type ObjPayload struct {
	attr []string
}

// Object returns a new ObjPayload.
func Object(arr ...string) *ObjPayload {
	obj := &ObjPayload{attr: arr}
	return obj
}

// Append adds member to ObjPayload.
func (p *ObjPayload) Append(arr ...string) *ObjPayload {
	p.attr = append(p.attr, arr...)
	return p
}

// String returns string represetation of ObjPayload.
func (p *ObjPayload) String() string {
	return fmt.Sprintf("{%s}", strings.Join(p.attr, ","))
}

// Any accepts any type fits Stringer interface.
func Any(key string, val fmt.Stringer) string {
	return fmt.Sprintf(`"%s": %s`, key, val.String())
}

// String generates member which has stirng as value.
func String(key, val string) string {
	return fmt.Sprintf(`"%s": "%s"`, key, val)
}

// Int generates member which has int as value.
func Int(key string, val int) string {
	return fmt.Sprintf(`"%s": %v`, key, val)
}

// Int32 generates member which has int32 as value.
func Int32(key string, val int32) string {
	return fmt.Sprintf(`"%s": %v`, key, val)
}

// Int64 generates member which has int64 as value.
func Int64(key string, val int64) string {
	return fmt.Sprintf(`"%s": %v`, key, val)
}

// Float32 generates member which has Float32 as value.
func Float32(key string, val float32) string {
	return fmt.Sprintf(`"%s": %v`, key, val)
}

// Float64 generates member which has float64 as value.
func Float64(key string, val float64) string {
	return fmt.Sprintf(`"%s": %v`, key, val)
}

// Attrs accepts several strings and stores it.
func Attrs(key string, arr ...string) string {
	return fmt.Sprintf(`"%s": {%s}`, key, strings.Join(arr, ","))
}

// ArrPayload stores values.
type ArrPayload struct {
	eles []string
}

// String returns string represetation of ArrPayload.
func (a *ArrPayload) String() string {
	return fmt.Sprintf("[%s]", strings.Join(a.eles, ","))
}

// Array stores values and returns a new ArrPayload.
func Array(arr ...string) *ArrPayload {
	ret := &ArrPayload{}
	for _, v := range arr {
		s := fmt.Sprintf(`"%s"`, v)
		ret.eles = append(ret.eles, s)
	}
	return ret
}

// AppendString appends stirngs as its values.
func (a *ArrPayload) AppendString(s ...string) *ArrPayload {
	for _, v := range s {
		str := fmt.Sprintf(`"%s"`, v)
		a.eles = append(a.eles, str)
	}
	return a
}

// AppendAny appends Stringers as its values.
func (a *ArrPayload) AppendAny(s ...fmt.Stringer) *ArrPayload {
	for _, v := range s {
		str := fmt.Sprintf("%s", v.String())
		a.eles = append(a.eles, str)
	}
	return a
}

// AppendInt appends ints as its values.
func (a *ArrPayload) AppendInt(arr ...int) *ArrPayload {
	for _, v := range arr {
		str := fmt.Sprintf("%+v", v)
		a.eles = append(a.eles, str)
	}
	return a
}
