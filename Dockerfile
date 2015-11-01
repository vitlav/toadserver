FROM quay.io/eris/base
MAINTAINER Eris Industries <support@erisindustries.com>

#TODO not be hardcoded
ENV MINTX_NODE_ADDR	http://0.0.0.0:46657/
ENV MINTX_SIGN_ADDR	http://0.0.0.0.4767
ENV MINTX_CHAINID	toadserver
ENV MINTX_PUBKEY	1C6A8E715DB881CD19829B2D438C3D90771FEC5CA5E77FA8136AE6D2B3EFBD06

ENV NAME         toadserver
ENV REPO 	 eris-ltd/$NAME
ENV BRANCH       master
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
