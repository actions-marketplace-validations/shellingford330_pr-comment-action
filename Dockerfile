FROM golang:1.18

WORKDIR /pr_comment

COPY . /pr_comment

RUN CGO_ENABLED=0 go build -o app .

ENTRYPOINT [ "/pr_comment/app" ]