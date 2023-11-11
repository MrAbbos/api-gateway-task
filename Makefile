CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --merge

copy-proto-module:
	rm -rf ${CURRENT_DIR}/protos
	rsync -rv --exclude=.git ${CURRENT_DIR}/ttu_protos/* ${CURRENT_DIR}/protos

gen-proto-module:
	./scripts/gen_proto.sh ${CURRENT_DIR}

rm-proto-omit-empty:
	chmod 744 ./scripts/rm_omit_empty.sh && ./scripts/rm_omit_empty.sh ${CURRENT_DIR}

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

build-image:
	docker build --no-cache --rm -t ${REGISTRY}/${GROUP}/${PROJECT_NAME}:${TAG} .
	docker tag ${REGISTRY}/${GROUP}/${PROJECT_NAME}:${TAG} ${REGISTRY}/${GROUP}/${PROJECT_NAME}:${ENV_TAG}

push-image:
	docker push ${REGISTRY}/${GROUP}/${PROJECT_NAME}:${TAG}
	docker push ${REGISTRY}/${GROUP}/${PROJECT_NAME}:${ENV_TAG}

swag-init:
	swag init -g api/api.go -o api/docs
	
run:
	go run cmd/main.go

linter:
	golangci-lint run