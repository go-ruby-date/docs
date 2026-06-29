# Roadmap

`go-ruby-date/date` is grown **test-first**, each capability differential-tested against MRI
rather than built in isolation. Ruby's `Date` and `DateTime` — the
deterministic, interpreter-independent slice extracted from rbgo's internals — is
**complete**.

| Stage | What | Status |
| --- | --- | --- |
| Julian-day model | Dates represented by astronomical Julian day number, the same internal model Ruby's `Date` uses, so arithmetic and conversions match exactly. | **Done** |
| Julian→Gregorian reform | The configurable reform date (default Italian 1582-10-15), with the missing-days gap handled the way reference Ruby handles it. | **Done** |
| `strftime` formatting | The full directive set (`%Y %m %d %H %M %S %j %A %a %B %b %z` …) producing byte-identical output to MRI's `Date#strftime`. | **Done** |
| `parse` & `strptime` | Heuristic `Date.parse` and format-driven `Date.strptime`, accepting the same inputs and raising on the same malformed ones as reference Ruby. | **Done** |
| Week-date & ordinal arithmetic | ISO week-date (`%G`/`%V`), ordinal (day-of-year) construction, and day/month arithmetic with Ruby's rollover semantics. | **Done** |
| Differential oracle & coverage | A wide date corpus formatted and parsed both here and by the system `ruby`/`date`, compared byte-for-byte; 100% coverage, gofmt + go vet clean, green across all six 64-bit Go arches and three OSes. | **Done** |

## Documented out-of-scope boundaries

These are **deliberate**, recorded so the module's surface is unambiguous:

- **No interpreter.** The library implements the deterministic algorithm; it
  never runs arbitrary Ruby. Anything that needs a live binding or evaluation is
  the consumer's job — that is why `rbgo` binds this module rather than the
  reverse.
- **Reference is reference Ruby (MRI).** Byte-for-byte conformance targets MRI's
  behaviour; differences across MRI releases are matched to the reference used by
  the differential oracle.
- **Standalone & reusable.** The module has no dependency on the Ruby runtime;
  the dependency runs the other way.

See [Usage & API](api.md) for the surface and [Why pure Go](why.md) for the
deterministic/interpreter split.
