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

`go get -u -v github.com/xnum/zkjson`

# Array

```go
func main() {
	arrName := zkjson.Array("Alice", "Bob") // Array only accepts string when creating.
	arrName.AppendString("Kate", "John")

	arrScore := zkjson.Array()
	arrScore.AppendInt(55, 66, 77, 88) // But you can append int.

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
{"name": ["Alice","Bob","Kate","John"],"score": [70,80,90,100]}
```

# object/member

```go
func main() {
	// Object accepts key-value pair.
	ezObj := zkjson.Object(
	    zkjson.String("Url", "http://www.google.com"),
	    zkjson.Int("TTL", 128),
	    zkjson.Float64("Delay", 3.21),
	)

	fmt.Println(ezObj)
}
```

output:
```
{"Url": "http://www.google.com","TTL": 128,"Delay": 3.21}
```

# complex struct

```go
	func main() {
	// Easy to consturct a complex structure.
	res := zkjson.Object(
		zkjson.String("name", "xnum"),
		zkjson.Int("age", 27),
		zkjson.Any("skill",
			zkjson.Array("C/C++", "Golang", "PHP"),
		), // Any() => `{..,"skill": ["C/C++", "Golang", "PHP"],..}`.
		json.Any("no idea", json.Object()), 
		zkjson.Attrs("education",
			zkjson.String("NCTU", "MS"),
			zkjson.String("NTCU", "BS"),
		),
	)

	fmt.Println(res)
}
```

```
{"name": "xnum","age": 27,"skill": ["C/C++","Golang","PHP"],"no idea": {},"education": {"NCTU": "MS","NTCU": "BS"}}
```
