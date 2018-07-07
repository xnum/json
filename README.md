# zkjson

zero knowledge json(zkjson) is a very simple library that only provides marshalling.

When we calls external service, we usually don't cares its request structure.

For example:

```
{
	"user_id": "xnum",
	"token": "27364107346120846143",
	"action": "buy",
	"req": [{
			"item": "apple",
			"amount": 100
		},
		{
			"item": "bear",
			"amount": 200
		},
		{
			"item": "charlie",
			"amount": 300
		}
	]
}
```

In Golang, you have to declare a structure with json tag to marshal json. So you are writing...

```
type req struct {
	UserID string `json:"user_id"`
	Token string `json:"token"`
	Action string `json:"buy"`
	Reqs []Orders `json:"req"`
}

req := &req{
	UserID: "xnum",
	....
}
```

I don't wanna write type, struct to deal with these dynamic json anymore.

Lots of JSON lib focus on how to unmarshal or parse and nobody provides an easy to use approach. So there is the solution.

# Install

`go get -u -v github.com/xnum/json`

# Array

```go
func main() {
	arrName := zkjson.Array("Alice", "Bob")
	arrName.Append("Kate", "John")

	arrScore := zkjson.Array()
	arrScore.Append(55, 66, 77, 88)

	students := zkjson.Object( // Using `Object` to combine it.
	    zkjson.Any("name", arrName),
	    zkjson.Any("score", arrScore),
	)

	fmt.Println(arrName)
	fmt.Println(arrScore)
	fmt.Println(students)
}
```

output:

```
["Alice","Bob","Kate","John"]
[70,80,90,100]
{"name":["Alice","Bob","Kate","John"],"score":[70,80,90,100]}
```

# object/member

```go
func main() {
	// Object accepts key-value pair.
	ezObj := zkjson.Object(
	    zkjson.Attr("Url", "http://www.google.com"),
	    zkjson.Attr("TTL", 128),
	    zkjson.Attr("Delay", 3.21),
	)

	fmt.Println(ezObj)
}
```

output:
```
{"Url":"http://www.google.com","TTL":128,"Delay":3.21}
```

# complex struct

```go
func main() {
	res := zkjson.Object(
		zkjson.Attr("name", "xnum"),
		zkjson.Attr("age", 27),
		zkjson.Attr("skill",
			zkjson.Array("C/C++", "Golang", "PHP"),
		),
		zkjson.Attr("no idea", zkjson.Object()), 
		zkjson.Attr("education", zkjson.Object(
			zkjson.Attr("NCTU", "MS"),
			zkjson.Attr("NTCU", "BS"),
		)),
	)

	fmt.Println(res)
}
```

```
{"name": "xnum","age": 27,"skill": ["C/C++","Golang","PHP"],"no idea": {},"education": {"NCTU": "MS","NTCU": "BS"}}
```

