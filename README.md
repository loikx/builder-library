# builder-library
Very small library for building golang object. 

--------------------------------------
Example of usage with exported fields:
```go
type Person struct {
	Name string
	Age  int
	Position  bool

	isNew bool
}

func main() {
	builderObject := builder.NewBuilder[Person]()
	builderObject.
		With("Name", "Leva").
		With("Age", 150).
		With("Position", true)

	person, errors := builderObject.Build()

	fmt.Println(person, errors)
}
```

Output:
```
&{Leva 150 true false} []
```
----------
Example of usage with non exported fields:
```go
type Person struct {
	Name string
	Age  int
	Position  bool

	isNew bool
}

func main() {
	builderObject := builder.NewBuilder[Person]()
	builderObject.
		With("Name", "Leva").
		With("Age", 150).
		With("Position", true).
		With("isNew", false)

	person, errors := builderObject.Build()

	fmt.Println(person, errors)
}
```

Output:
```
&{Leva 150 true false} [can not set field isNew with type bool]
```

At this case we will get an array of errors.

