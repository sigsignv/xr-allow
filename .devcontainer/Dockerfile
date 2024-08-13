FROM golang:1.22-bookworm AS base

FROM base AS build

RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/cweill/gotests/gotests@v1.6.0
RUN go install github.com/fatih/gomodifytags@v1.16.0
RUN go install github.com/josharian/impl@v1.1.0
RUN go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

FROM base AS final

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=${USER_UID}

ENV LANG=C.UTF-8

RUN groupadd -g ${USER_GID} ${USERNAME} \
 && useradd -u ${USER_UID} -g ${USER_GID} -G adm,sudo -m -s /bin/bash ${USERNAME}

RUN --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    export DEBIAN_FRONTEND=noninteractive && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
        bash-completion \
        bind9-dnsutils \
        less \
        publicsuffix \
        sudo \
        vim-nox \
        whiptail \
        zip

RUN chown -R ${USER_UID}:${USER_GID} /go

COPY --from=build --chown=${USER_UID}:${USER_GID} /go/bin /go/bin

RUN echo "${USERNAME} ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/user \
 && chmod 440 /etc/sudoers.d/user

USER ${USERNAME}
CMD [ "/bin/bash", "-l" ]
