# strongword - Password strength validator for GO
<p>
  <a href="https://pkg.go.dev/github.com/bounoable/strongword">
    <img alt="GoDoc" src="https://img.shields.io/badge/godoc-reference-purple">
  </a>
  <img alt="Version" src="https://img.shields.io/github/v/tag/bounoable/strongword" />
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

Simple utility library to validate password strengths against a set of rules.

## Install

```sh
go get github.com/bounoable/strongword
```

## Usage

### Default rule set

```go
import "github.com/bounoable/strongword"

errs := strongword.Validate("weakpassword")

for _, err := range errs {
  fmt.Println(err)
}
```

### Custom rule set

```go
import "github.com/bounoable/strongword"

errs := strongword.Validate(
  "weakpassword",
  strongword.MinLength(12),
  strongword.SpecialChars(3),
  strongword.Regexp(regexp.MustCompile("(?)[0-9]{4}"))
)

for _, err := range errs {
  fmt.Println(err)
}
```
