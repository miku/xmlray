SHELL := /bin/bash
TARGETS = xmlray

# http://docs.travis-ci.com/user/languages/go/#Default-Test-Script
test:
	go get -d && go test -v

imports:
	goimports -w .

fmt:
	go fmt ./...

all: fmt test
	go build

install:
	go install

clean:
	go clean
	rm -f coverage.out
	rm -f $(TARGETS)
	rm -f xmlray-*.x86_64.rpm
	rm -f packaging/debian/xmlcutty_*.deb
	rm -f xmlcutty_*.deb
	rm -rf packaging/debian/xmlray/usr

cover:
	go get -d && go test -v	-coverprofile=coverage.out
	go tool cover -html=coverage.out

xmlray: cmd/xmlray/main.go visitors.go
	go build -o xmlray cmd/xmlray/main.go

# ==== packaging

# deb: $(TARGETS)
# 	mkdir -p packaging/debian/xmlray/usr/sbin
# 	cp $(TARGETS) packaging/debian/xmlray/usr/sbin
# 	cd packaging/debian && fakeroot dpkg-deb --build xmlray .
# 	mv packaging/debian/xmlray*deb .

# rpm: $(TARGETS)
# 	mkdir -p $(HOME)/rpmbuild/{BUILD,SOURCES,SPECS,RPMS}
# 	cp ./packaging/rpm/xmlray.spec $(HOME)/rpmbuild/SPECS
# 	cp $(TARGETS) $(HOME)/rpmbuild/BUILD
# 	./packaging/rpm/buildrpm.sh xmlray
# 	cp $(HOME)/rpmbuild/RPMS/x86_64/xmlray*.rpm .
