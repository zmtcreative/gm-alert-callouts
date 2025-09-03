module github.com/zmtcreative/gm-alert-callouts

go 1.23.0

toolchain go1.24.5

retract [v0.0.0, v0.7.0]

require github.com/yuin/goldmark v1.7.13

require (
	github.com/jeandeaual/go-locale v0.0.0-20250612000132-0ef82f21eade
	golang.org/x/text v0.27.0
)

require golang.org/x/sys v0.27.0 // indirect
