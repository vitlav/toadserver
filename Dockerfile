FROM quay.io/eris/base
MAINTAINER Eris Industries <support@erisindustries.com>

ENV NAME         toadserver
ENV REPO 	 eris-ltd/$NAME
ENV BRANCH       develop
ENV BINARY_PATH  $NAME
ENV CLONE_PATH   $GOPATH/src/github.com/$REPO
ENV INSTALL_PATH $INSTALL_BASE/$NAME

WORKDIR $CLONE_PATH
RUN git clone -q https://github.com/$REPO $CLONE_PATH && \
 git checkout -q $BRANCH && \
 go install

RUN rm -rf $GOPATH/src/* && \
  unset NAME && \
  unset INSTALL_BASE && \
  unset REPO && \
  unset CLONE_PATH && \
  unset BINARY_PATH && \
  unset INSTALL_PATH && \
  unset BRANCH

USER $USER
WORKDIR $ERIS

VOLUME $ERIS
EXPOSE 11113
CMD ["toadserver"]
