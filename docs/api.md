# Usage & API

The public API lives at the module root (`github.com/go-ruby-date/date`). It is **Ruby-shaped but Go-idiomatic**: `Parse` / `Strptime` / `Strftime` mirror Ruby's `Date.parse` / `Date.strptime` / `Date#strftime`, while the surface follows Go conventions — an explicit `error`, value types, no global state.

!!! success "Status: implemented"
    The library is built and importable as `github.com/go-ruby-date/date`, bound into
    `rbgo` as a native module; see [Roadmap](roadmap.md).

## Install

```sh
go get github.com/go-ruby-date/date
```

## Worked example

```go
d, _ := date.Parse("2026-06-29")
s := d.Strftime("%A, %d %B %Y")     // "Monday, 29 June 2026"
d2, _ := date.Strptime("29/06/2026", "%d/%m/%Y")
```

## Shape

```go
// Parse heuristically parses a date string (Date.parse).
func Parse(s string) (Date, error)

// Strptime parses s according to a strftime-style format (Date.strptime).
func Strptime(s, fmt string) (Date, error)

// Strftime formats the date with the strftime directive set (Date#strftime).
func (d Date) Strftime(fmt string) string

// JD returns the astronomical Julian day number.
func (d Date) JD() int
```

## MRI conformance

Correctness is defined by reference Ruby. A **differential oracle** runs a wide
corpus through both the system `ruby` and this library and compares the results
**byte-for-byte** — not approximated from memory. The oracle tests skip
themselves where `ruby` is not on `PATH` (e.g. the qemu arch lanes), so the
cross-arch builds still validate the library.

## Relationship to Ruby

`go-ruby-date/date` is **standalone and reusable**, and is the backend bound into
[go-embedded-ruby](https://github.com/go-embedded-ruby/ruby) by `rbgo` as a
native module — the same way [go-ruby-regexp](https://github.com/go-ruby-regexp)
and [go-ruby-erb](https://github.com/go-ruby-erb) are bound. The dependency runs
the other way: this library has no dependency on the Ruby runtime.
