.PHONY: rsakey

rsakey:
	rm ./setting/key ./setting/key.pub >> /dev/null; exit 0
	ssh-keygen -t rsa -b 512 -C "git.jediautocare.com" -f ./setting/key -N "" -q

fmt:
	go fmt ./...
	goimports -w $$(find . -type f -name '*.go')

check:
	make fmt
	golint -set_exit_status ./...
	golangci-lint run ./...

init:
	# git push check
	echo "#!/bin/sh\nmake check" > ./.git/hooks/pre-push
	chmod +x ./.git/hooks/pre-push
	cd && go get -u golang.org/x/tools/cmd/goimports
	cd && go get -u golang.org/x/lint/golint
	cd && go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.1