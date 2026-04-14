# Go Range Loop Pitfalls

This folder contains small runnable examples of common `for range` pitfalls in Go and the correct patterns to use instead.

## Run an example

From this folder:

```bash
go run ./01_address_of_range_var
go run ./02_modify_copy_not_source
go run ./03_map_value_copy
go run ./04_append_during_range
go run ./05_goroutine_capture_pre122
```

## What each example shows

1. `01_address_of_range_var`: taking `&v` vs `&slice[i]`.
2. `02_modify_copy_not_source`: mutating range value copy vs mutating original element.
3. `03_map_value_copy`: map value updates require write-back (`m[k] = v`).
4. `04_append_during_range`: range over a slice uses length snapshot at loop start.
5. `05_goroutine_capture_pre122`: closure capture behavior and modern safe pattern.
