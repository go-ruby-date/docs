// SPDX-License-Identifier: BSD-3-Clause
package main

import "github.com/go-ruby-date/date"

func main() {
	s := "2024-03-15"
	d, _ := date.Parse(s, true)
	fmt := "%Y-%m-%dT%H:%M:%S %A"
	bench("parse-iso", 2000, func() { v, _ := date.Parse(s, true); sink = v })
	bench("strftime", 2000, func() { sink = d.Strftime(fmt) })
	bench("plus-days", 5000, func() { sink = d.Plus(1) })
}
