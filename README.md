# Go test coverage action

Use this action to run tests in Golang and make sure are complying with the coverage. Support the use of `go test` and `ginkgo` commands

## Usage

Basic usage using `go test`:

```yaml
- name: Coverage test
  uses: joyarzun/go-testcoverage-action@v1
  with:
    coverage-threshold: 80
```

Basic usage using `ginkgo`:

```yaml
- name: Coverage test
  uses: joyarzun/go-testcoverage-action@v1
  with:
    use-ginkgo: true
    coverage-threshold: 80
```

## Usage options

```yaml
install-go:
  description: |
    By default the action install a version of Go appropriate for building and running.
    If set to false, you will need to have installed Go.
  required: true
  default: true
go-version:
  description: |
    Version of Go
  required: false
  default: 1.19.x
cache-key:
  description: |
    Custom string to include in the cache key. This is useful when using multiple Go versions.
  required: false
coverage-threshold:
  description: |
    The threshold in percentage to consider the test successful. Default is 0 so any test will pass
  required: true
  default: 0
use-ginkgo:
  description: |
    Use ginkgo to run tests and get the coverage
  required: false
  default: false
workdir:
  description: |
    Go base directory
  required: false
  default: .
```
