
GLADES=$(wildcard ui/*.glade)
GOS=$(wildcard *.go)

all: dnetgtk

ui-bindata.go: $(GLADES)
	-rm -v ui/*~
	go-bindata -o ui-bindata.go ui


dnetgtk: $(GOS) ui-bindata.go Makefile
	go build -v


