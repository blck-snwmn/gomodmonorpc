# gomodmonorpc 

gomodmonorpc is go multi module, modular monolith(use Protocol Buffers) sample.

## Run
```bash
go run cmd/server/main.go
```

## Workspace mode

1. Run command in root directory.
   ```
   go work init  . modules/order
   ```

2. Then, `go.work` file will be generated.
   ```
   go 1.19

   use (
       ./main
       ./modules/order
   )
   ```
