# FROM golang:1.15.3-alpine3.12 AS build

# RUN apk --update --no-cache add gcc g++ make ca-certificates git tzdata openssh
# ENV TZ=Asia/Almaty

# ARG SSH_PRIVATE_KEY
# ARG PRIVATE_PKG_REPOSITORY
# RUN mkdir /root/.ssh/ && \
#     echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa && \
#     chmod 600 /root/.ssh/id_rsa && \
#     touch /root/.ssh/known_hosts && \
#     ssh-keyscan ${PRIVATE_PKG_REPOSITORY} >> /root/.ssh/known_hosts

# WORKDIR /go/src/project
# COPY . .
# RUN go env -w GOPRIVATE=${PRIVATE_PKG_REPOSITORY} && \
#     git config --global url."git@${PRIVATE_PKG_REPOSITORY}:".insteadOf "https://${PRIVATE_PKG_REPOSITORY}/" &&\
#     go mod download && \
#     go build -o /go/bin/app main.go

FROM alpine:3.12

RUN apk --update --no-cache add tzdata
ENV TZ=Asia/Almaty

WORKDIR /usr/bin
# COPY --from=build /go/bin/app .

EXPOSE 3001

ENTRYPOINT ["./app"]

CMD ["serve"]
