# frozen_string_literal: true
# SPDX-License-Identifier: BSD-3-Clause
require "date"
require_relative "_harness"
s   = "2024-03-15"
d   = Date.parse(s)
fmt = "%Y-%m-%dT%H:%M:%S %A"
bench("parse-iso",   2000) { Date.parse(s) }
bench("strftime",    2000) { d.strftime(fmt) }
bench("plus-days",   5000) { d + 1 }
