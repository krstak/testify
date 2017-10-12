# Testify
Small assertion library

## Install
```
go get github.com/krstak/testify
```

## Usage

```go
func TestTestify(t *testing.T) {
    testify.Equal("super", "super")(t)
    testify.NotEqual("car", "street")(t)
    testify.Nil(nil)(t)
    testify.NotNil("house")(t)
    testify.True(10 > 2)(t)
    testify.False(10 == 2)(t)
}
```

### Important
Do not forget to call `(t)` function at the end of assertion.