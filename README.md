#### Wat
Dumb utility to migrate Godeps deps into ./vendor using gvt.

#### Use

	go build -o godep-to-gvt main.go
	cp godep-to-gvt $GOPATH/{path to your project}
	cd $GOPATH/{path to your project}
	./godep-to-gvt -fp=`pwd` -latest

If you omit the `-latest` flag, the revision specified in the Godeps.json file will be fetched.
