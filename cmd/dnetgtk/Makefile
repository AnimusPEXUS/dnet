
GLADES=$(wildcard ui/*.glade)
GOS=$(wildcard *.go)

all: dnetgtk

bindata.go: $(GLADES)
	-rm -v ui/*~
	go-bindata ui


dnetgtk: $(GOS)
	go build


