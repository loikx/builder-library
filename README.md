# builder-library
Very small library for building golang object. 
Example of usage:
```go
type Person struct {
	Name string
	Age  int
	Lox  bool

	isNew bool
}

func main() {
	builderObject := builder.NewBuilder[Person]()
	builderObject.
		With("Name", "Leva").
		With("Age", 150).
		With("Lox", true)

	person, err := builderObject.Build()

	fmt.Println(person, err)
}

```
