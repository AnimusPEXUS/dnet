
# This Makefile is needed only to rebuild ui_bindata.go files.
# Only developers of dnetgtk may need this.
# Users should be able simply go get dnetgtk without any additional
# movements


all:
	find -type f -name '*.glade~' -print -delete
	python3 ./build_ui.py
	go build -v
