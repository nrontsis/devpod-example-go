FROM golang:1.23

RUN apt-get -q update && \
    apt-get install -y --no-install-recommends \
    sudo \
    && rm -rf /var/lib/apt/lists/*
asldkfjlaskdj
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# Pre-create files to allow devpod to run without root access
RUN mkdir /var/run/devpod && chown -R $USER_UID:$USER_UID /var/run/devpod
RUN mkdir /var/devpod && touch /var/devpod/chownWorkspace.marker && chown -R $USER_UID:$USER_UID /var/devpod
RUN [ ! -d /home/root ] && ln -s /root /home/root
USER $USER_UID
