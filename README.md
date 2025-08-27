# Test for go-git issue #647

This is a test for [go-git issue #647](https://github.com/go-git/go-git/issues/647).

The server in [main.go](./main.go) serves the repository of the current working directory on <http://localhost:8080>.

To reproduce the error, you can either run the test with `go test ./...` or manually as follows:

1. Start the git HTTP server:

   ```sh
   go run main.go
   ```

2. In a separate terminal, run the following:

   ```sh
   git clone --no-progress http://localhost:8080
   ```

3. You should see an output similar to this:

   ```text
   Cloning into 'localhost'...
   fatal: protocol error: bad line length character: PACK
   fatal: protocol error: bad pack header
   ```
