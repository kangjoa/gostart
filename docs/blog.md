# Error Handling in Go

_Like a guardrail, you won’t notice Go’s error handling until you need it._

![Guardrail on Pacific Coast Highway](pch-guardrail.jpeg)

I’m most familiar with error handling using [try...catch](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/try...catch) blocks in JavaScript. With `try…catch blocks`, you try the success path first and if that doesn’t work, it throws, and we jump to the catch block.

In the JavaScript example below, we try to create a file from a template but if something goes wrong, we catch the error and log it. However, JavaScript doesn’t require us to include a `try…catch` block. You could simply only include the code inside the try block alone, ignore that an error might happen, and take your chances, risking a crash.

```js
try {
  createFromTemplate(
    'templates/go.mod.template',
    filepath.Join(projectName, 'go.mod'),
    projectName,
  );
} catch (error) {
  console.error(error);
}
```

Alternatively, Go has explicit error handling. This means Go requires you to make a choice about every error, either to check, return, or discard them.

In the Go example below, we immediately face the possibility of failure and deal with it. It’s like slipping off the road and clipping the guardrail, but the guardrail keeps us from going off the cliff.

```go
if err := createFromTemplate("templates/go.mod.template",
	filepath.Join(projectName, "go.mod"), projectName); err != nil {
	fmt.Printf("Error creating go.mod: %v\n", err)
	return
}
```

You can’t ignore the error path automatically, but to deliberately discard the error value, you can use the blank identifier, an underscore (\_). This means if an error occurs, your program will just ignore the failure by not reacting to it. [Discarding](https://go.dev/doc/effective_go#blank) is highly discouraged because if the template does not exist, the error is ignored and the code may later fail or panic. Imagine driving on a cliff without a guardrail: while you’re squarely on the road you won’t notice the need for a guardrail, but if you slip off the road, you’ll really wish you had one.

```go
// Discarding is NOT recommended
_ = createFromTemplate("templates/go.mod.template",
	filepath.Join(projectName, "go.mod"), projectName)
```

## Panic

While errors are expected failure paths like missing files or bad inputs, `panic` is a built-in function that stops a program. In the example below, the function assumes b != 0, but if you violate that then the program stops immediately.

```go
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}

func main() {
	result := divide(10, 0) // panic happens
}

```

If you ran this, you would see something like “panic: division by zero” in the terminal.

Panics should be sparingly. Instead of using `panic` above, returning an error would be more appropriate because division by zero is bad input, not an unrecoverable bug.

```go
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
```

One example of where `panic` would be appropriate is when a package initialization fails. If a package can’t initialize, then we can’t continue execution. In this case, it could be better to fail very quickly with a `panic`.

```go
func init() {
	if err := loadCriticalResource(); err != nil {
		panic(fmt.Sprintf("failed to load critical resource: %v", err))
	}
}
```

Go’s approach to error handling might feel repetitive, but it’s just a truthful reminder that errors can always happen and you should prepare for them (and that’s why we have guardrails!).
