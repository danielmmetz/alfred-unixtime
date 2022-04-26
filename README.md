# alfred-unixtime

A tiny utility intended for use in an alfred plugin.

Supports:

`alfred-unixtime`: returns the current unix time in seconds

`alfred-unixtime now`: returns the current unix time in seconds

`alfred-unixtime $n`: translates the input time, assumed to be in either unix seconds or unix ms, to a local, readable datetime.

## Caveats

Very likely only works on arm64 due to embedding a go binary in the workflow.
