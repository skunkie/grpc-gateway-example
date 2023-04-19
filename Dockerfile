FROM golang:alpine

WORKDIR /build

COPY --from=swaggerapi/swagger-ui /usr/share/nginx/html/ /swagger-ui

RUN apk update \
    && apk add --no-cache curl git jq make protobuf-dev \
    && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2 \
    && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2 \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 \
    && go install github.com/envoyproxy/protoc-gen-validate@v0.10.1 \
    && go install github.com/go-bindata/go-bindata/...@latest \
    && git clone --depth 1 https://github.com/googleapis/googleapis.git ${GOPATH}/google/googleapis \
    && download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
    jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url') \
    && curl -o /usr/local/bin/swagger -L'#' "$download_url" \
    && chmod +x /usr/local/bin/swagger \
    && download_url=$(curl -s https://api.github.com/repos/pseudomuto/protoc-gen-doc/releases/latest | \
    jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url') \
    && curl -L'#' "$download_url" | tar xvz -f - -C /usr/local/bin protoc-gen-doc

CMD protoc \
        --proto_path ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.15.2 \
        --proto_path ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.10.1 \
        --proto_path ${GOPATH}/google/googleapis \
        --proto_path ./proto \
        --go-grpc_out ${out_path} \
        --go_out  ${out_path} \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_out ${out_path} \
        --openapiv2_opt allow_merge=true \
        --openapiv2_opt disable_default_responses=true \
        --openapiv2_opt disable_service_tags=true \
        --openapiv2_opt logtostderr=true \
        --openapiv2_opt merge_file_name=api \
        --openapiv2_out ${out_path} \
        --validate_out="lang=go:${out_path}" \
        --experimental_allow_proto3_optional \
        --plugin=protoc-gen-doc=/usr/local/bin/protoc-gen-doc \
        --doc_out=./docs \
        --doc_opt=markdown,proto.md \
        ./proto/*.proto \
    && mv ${out_path}/api.swagger.json /swagger-ui/swagger.json \
    && sed -i "s|https://petstore.swagger.io/v2/swagger.json|./swagger.json|g" /swagger-ui/swagger-initializer.js \
    && sed -i '/deepLinking: true,/a defaultModelsExpandDepth: 0,' /swagger-ui/swagger-initializer.js \
    && sed -i '/deepLinking: true,/a validatorUrl: null,' /swagger-ui/swagger-initializer.js \
    && go-bindata -o ${out_path}/swagger/swagger.go -nomemcopy -pkg=swagger -prefix=/swagger-ui/ /swagger-ui \
    && mkdir -p ${GOPATH}/src/${module_name}/pkg/httpclient \
    && swagger generate client --spec=/swagger-ui/swagger.json --target=${GOPATH}/src/${module_name}/pkg/httpclient \
    && find ./pkg/httpclient -mindepth 1 -delete && cp -R ${GOPATH}/src/${module_name}/pkg/httpclient ./pkg
