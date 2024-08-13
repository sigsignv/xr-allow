FROM golang:1.22-bookworm

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=${USER_UID}

ENV LANG=C.UTF-8

RUN groupadd -g ${USER_GID} ${USERNAME} \
 && useradd -u ${USER_UID} -g ${USER_GID} -G adm,sudo -m -s /bin/bash ${USERNAME}

RUN export DEBIAN_FRONTEND=noninteractive \
 && apt-get update \
 && apt-get upgrade -y \
 && apt-get install -y --no-install-recommends \
        bash-completion \
        bind9-dnsutils \
        less \
        publicsuffix \
        sudo \
        vim-nox \
        whiptail \
        zip \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

RUN echo "${USERNAME} ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/user \
 && chmod 440 /etc/sudoers.d/user

USER ${USERNAME}
CMD [ "/bin/bash", "-l" ]
