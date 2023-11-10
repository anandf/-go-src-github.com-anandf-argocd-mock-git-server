FROM golang:1.21 as builder

WORKDIR /

COPY . /

RUN CGO_ENABLED=0 go build -o argocd-git-mock-server main.go


FROM scratch

COPY --from=builder argocd-git-mock-server /

ENTRYPOINT [ "/argocd-git-mock-server" ]
