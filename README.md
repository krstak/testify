# Testify
small assertion library

## Install
```
go get github.com/krstak/testify
```

## Usage

```go
func TestTestify(t *testing.T) {
    assert := testify.New(t)

    assert.Equal("super", "super")
    assert.NotEqual("car", "street")
    assert.Nil(nil)
    assert.NotNil("house")
    assert.True(10 > 2)
    assert.False(10 == 2)
}
```