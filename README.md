pagerduty
=========
View upcoming oncall schedules through REST API

Build
=====
go build .

Run
===
You need to set these environment variables, then run the pagerduty binary
```
PAGERDUTY_KEY=xxxxxxxxxxxxxxxxxxxx
PAGERDUTY_PORT=xxxx
PAGERDUTY_USER=xxxxxxx
```
Then, `curl localhost:$PAGERDUTY_PORT/oncalls`
