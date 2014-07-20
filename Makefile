all:
	go get -d -v ./...
	go get -u gopkg.in/check.v1
	go get code.google.com/p/go.tools/cmd/cover
	go get -u github.com/mattn/goveralls

clean:
	rm -rf path-matcher reports .gopack