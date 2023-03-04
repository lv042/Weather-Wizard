###############################################################################################################################
# This Dockerfile constructs a lightweight linux environment where a connection will be setup to Localhost Run
# Source: https://localhost.run/docs/cli
# 
# Author: Mats Otten
# Date: 31 july 2022
###############################################################################################################################

FROM alpine

RUN apk add --no-cache --update openssh
RUN apk add --no-cache --update openssh-keygen
RUN apk add --no-cache --update php
ADD localhost-run.sh /
RUN chmod +x /localhost-run.sh

RUN ssh-keygen -q -t rsa -b 2048 -q -N "" -f ${HOME}/.ssh/id_rsa

CMD sh /localhost-run.sh