GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY  = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
IMAGE_REPO = ryane

.PHONY: docker
docker:
	docker build --force-rm -t ${IMAGE_REPO}/sampleapp:${GIT_SHA} -f ./Dockerfile .

docker.push: docker
	docker push ${IMAGE_REPO}/sampleapp:${GIT_SHA}
	docker tag ${IMAGE_REPO}/sampleapp:${GIT_SHA} ${IMAGE_REPO}/sampleapp:kubecon
	docker push ${IMAGE_REPO}/sampleapp:kubecon
