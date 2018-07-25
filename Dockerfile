FROM nlepage/golang_wasm AS builder

COPY ./ src/app/
RUN go build -o test.wasm app

FROM nlepage/golang_wasm:nginx

COPY --from=builder /go/test.wasm /usr/share/nginx/html/
COPY index.html /usr/share/nginx/html/
