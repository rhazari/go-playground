# GoStringer Example

This example shows how `fmt.GoStringer` customizes debug output for `%#v`.

## Run

```bash
go run ./go-gostringer-example
```

## Expected behavior

- `%v` uses `String()`.
- `%#v` uses `GoString()`.
