[![Go Report Card](https://goreportcard.com/badge/github.com/JingIsCoding/number)](https://goreportcard.com/report/github.com/JingIsCoding/number)
[![Go Reference](https://pkg.go.dev/badge/github.com/JingIsCoding/number.svg)](https://pkg.go.dev/github.com/JingIsCoding/number)

<!-- Go number -->
## Go number
A convenient abstraction to deal with basic arithmetic operation across different number types and precisions

<!-- What problems we are trying to solve -->
### What problems we are trying to solve

##### 1.Type casting when doing arithmetic operation amoung different number types or precisions

>Instead of 
```go
  var a int = 1
  var b float32 = 1.0
  var c float64 = 2.0
  var d int64 = int64(float64((float32(a) + b)) * float64(c))
```
>You can do
```go
  var a int = 1
  var b float32 = 1.0
  var c float64 = 2.0
  var d = number.Of(a).Add(b).Multiply(c).GetInt()
```
##### 2.Error handing when divide by 0
In Golang, when integer number divide by 0 it will panic with error `runtime error: integer divide by zero`\
whereas float number divide by 0 will result in a value of +Inf or -Inf
   
Now we make any division operation returns an error of type `number.DivideByZeroError`
```go
  var a number.Number = number.Of(1.0)
  var b number.Number = number.Of(0.0)
  var d, err = a.Divide(b)
  if err != nil {
    log.Panicln(err)
  }
```
##### 3.Convienet method for common operation
```go
  var a number.Number = number.Of(0.5)
  var b number.Number = a.RoundUp()
  var c number.Number = a.RoundDown()
  var d number.Number = a.ShiftDecimal(1)

  log.Println(b.GetFloat(), c.GetFloat(), d.GetFloat()) // 1 0 5
```
## Installation
```
go get -u github.com/JingIsCoding/number
```

## Credits
Thanks for @xchen1189 for co-creating this project.

## Feedback
Feel free to submit feedback to the issue or submit a PR for bug fix or feature extension.



