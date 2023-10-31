VERSION ?= 1.0.0

release:
	go mod tidy
	git add .
	git commit -m "[RELEASE] GO-HONUA-LOGGER: changes for v${VERSION}"
	git tag v${VERSION}
	git push origin v${VERSION}
	git push