module github.com/f4rx/otus-golang/hw01_hello_otus

go 1.16

require golang.org/x/example v0.0.0-20210407023211-09c3a5e06b5d

// go mod edit -replace=golang.org/x/example@master=github.com/golang/example@master
replace golang.org/x/example v0.0.0-20210407023211-09c3a5e06b5d => github.com/golang/example v0.0.0-20210407023211-09c3a5e06b5d
