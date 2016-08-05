#### Wat
Dumb utility to migrate Godeps deps into ./vendor using gvt.

#### Use

	go get -u https://github.com/chris-skud/godep-to-gvt
	
At the root of your project:

	godep-to-gvt -fp=`pwd` -latest

If you omit the `-latest` flag, the revision specified in the Godeps.json file will be fetched.
