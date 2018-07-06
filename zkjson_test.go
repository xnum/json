package json

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type testJsonSuite struct {
	suite.Suite
}

func (s *testJsonSuite) TestObj() {
	{
		obj := Object()
		s.Require().Equal("{}", obj.String())
	}

	{
		obj := Object(`"a": 1`)
		s.Require().Equal(`{"a": 1}`, obj.String())
	}

	{
		obj := Object(String("a", "1"))
		s.Require().Equal(`{"a": "1"}`, obj.String())
	}

	{
		subObj := Object()
		obj := Object(Any("sub", subObj))
		s.Require().Equal(`{"sub": {}}`, obj.String())
	}

	{
		obj := Object(Int("a", 1), Int("b", 2))
		s.Require().Equal(`{"a": 1,"b": 2}`, obj.String())
	}

	{
		obj := Object(Int64("a", 1), Float64("b", 0.4))
		s.Require().Equal(`{"a": 1,"b": 0.400000}`, obj.String())
	}
}

func (s *testJsonSuite) TestObjAttrs() {
	{
		obj := Object(Attrs("a"))
		s.Require().Equal(`{"a": {}}`, obj.String())
	}

	{
		obj := Object(Attrs("Q", Int("a", 1), Int("b", 2)))
		s.Require().Equal(`{"Q": {"a": 1,"b": 2}}`, obj.String())
	}

	{
		obj := Object(Attrs("a", Attrs("b", Attrs("c", Attrs("d")))))
		s.Require().Equal(`{"a": {"b": {"c": {"d": {}}}}}`, obj.String())
	}
}

func (s *testJsonSuite) TestObjAppend() {
	{
		obj := Object(String("a", "A"), Int("b", 99), String("c", "C"))
		obj2 := Object()
		obj2.Append(String("a", "A"))
		obj2.Append(Int("b", 99))
		obj2.Append(String("c", "C"))

		s.Require().Equal(obj2.String(), obj.String())
	}
}

func (s *testJsonSuite) TestArray() {
	{
		arr := Array()
		s.Require().Equal(`[]`, arr.String())
	}

	{
		arr := Array("a", "b", "c")
		s.Require().Equal(`["a","b","c"]`, arr.String())
	}

	{
		arr := Array()
		arr.AppendString("a", "b")
		arr.AppendString("c")
		s.Require().Equal(`["a","b","c"]`, arr.String())
	}

	{
		arr := Array()
		arr.AppendString("a", "b")
		arr.AppendInt(3)
		s.Require().Equal(`["a","b",3]`, arr.String())
	}
}

func (s *testJsonSuite) TestArrayAppendAny() {
	{
		arr := Array()
		arr.AppendAny(Object(), Object(String("a", "b")))
		s.Require().Equal(`[{},{"a": "b"}]`, arr.String())
	}

	{
		arr := Array()
		arr.AppendAny(Object(), Array())
		s.Require().Equal(`[{},[]]`, arr.String())
	}
}

func (s *testJsonSuite) TestComplexObject() {
	user := Object()
	user.Append(String("user_id", "3345678"))
	user.Append(String("name", "xnum"))
	user.Append(String("country", "tw"))

	location := Array()
	location.AppendFloat64(124.012341, 23.998745)

	event := Object(
		String("action", "login"),
		String("service", "auth"),
	)

	// Once you put an anything as parameter, it has been marshalled.
	// Any modifications of object are not working.
	payload := Object(Any("user", user), Any("loc", location), Any("event", event))
	data := payload.String()

	// When you modifies event after it used, it doesn't change anything.
	event.Append(String("data", "gmail"))

	s.Require().Equal(data, payload.String())
}

func TestJson(t *testing.T) {
	suite.Run(t, &testJsonSuite{})
}
