# golang:1 is the tag for latest 1.x stable release
# using alpine variant only because "apk add docker" is easier than on debian
FROM golang:1.16.3-alpine3.13

# labels for github actions UI
LABEL "com.github.actions.name"="operator-sdk"
LABEL "com.github.actions.description"="operator-sdk image builder"
LABEL "com.github.actions.icon"="layers"
LABEL "com.github.actions.color"="red"

# add operator-sdk release binary
ARG operator_sdk_version=v1.5.0
ARG operator_sdk_base=https://github.com/operator-framework/operator-sdk/releases/download/
ARG sdk_cli=/usr/bin/operator-sdk

ADD ${operator_sdk_base}${operator_sdk_version}/operator-sdk_linux_amd64 ${sdk_cli}
RUN chmod +x ${sdk_cli}

# could also add "alpine-sdk" to apk add to enable the go tools from this image
RUN apk add --update docker
RUN apk add --update make
RUN go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.4.1

# usage: in GitHub Actions, execute "operator-sdk build image:version"
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]