build:
	cp -r app/ src/app
	go get -v github.com/spf13/cobra/cobra

backup:
	mv src/app app/

clean:
	mv src/app app/
	rm -rf pkg/ bin/ src/