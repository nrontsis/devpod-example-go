FROM mcr.microsoft.com/devcontainers/go:0-1.20-bullseye

USER root
# Pre-create files to allow devpod to run without root access
RUN mkdir /var/run/devpod && chown -R $USERNAME:$USERNAME /var/run/devpod
RUN mkdir /var/devpod && touch /var/devpod/chownWorkspace.marker && chown -R $USERNAME:$USERNAME /var/devpod
RUN [ ! -d /home/root ] && ln -s /root /home/root
USER vscode
