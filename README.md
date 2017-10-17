# Testify
Small assertion library

## Install
```
go get github.com/krstak/testify
```

## Usage

```go
func TestTestify(t *testing.T) {
    testify.Equal(t)("super", "super")
    testify.NotEqual(t)("car", "street")
    testify.Nil(t)(nil)
    testify.NotNil(t)("house")
    testify.True(t)(10 > 2)
    testify.False(t)(10 == 2)
}
```
