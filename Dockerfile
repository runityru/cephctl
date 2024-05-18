ARG CEPH_SOURCE_IMAGE=quay.io/ceph/ceph
ARG CEPH_VERSION=18.2.2

FROM ${CEPH_SOURCE_IMAGE}:v${CEPH_VERSION}

ADD --chown=root:root dist/cephctl_linux_amd64_v3/cephctl /usr/bin/cephctl
