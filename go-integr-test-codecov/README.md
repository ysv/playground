# Go Code Coverage for Integration tests

Based on https://go.dev/blog/integration-test-coverage

## Steps

1. Prepare folder for go coverage
```bash
rm -rf covdatafiles && mkdir covdatafiles
```

3. Start go HTTP server
```bash
GOCOVERDIR=covdatafiles go run -cover main.go
⇨ http server started on [::]:8080
```

4. Make HTTP request to cover new lines of code
```bash
curl localhost:8080/time
{"time":"2024-01-09T16:15:05Z"}
```

5. Check the coverage report. Note that coverage has not been reported yet.
```bash
go tool covdata percent -i=covdatafiles

	command-line-arguments		coverage: 0.0% of statements
```

6. Stop HTTP server by ctr+C

7. View the coverage report
```bash
go tool covdata percent -i=covdatafiles 

	command-line-arguments		coverage: 89.5% of statements
```

```bash
go tool covdata textfmt -i=covdatafiles -o=cov.txt
go tool cover -html=cov.txt
```