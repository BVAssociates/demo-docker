FROM i386/ubuntu:14.04
MAINTAINER Vincent Bauchart <bauchart@bvassociates.fr>

# Fixes Docker Automated Build problem
RUN ln -s -f /bin/true /usr/bin/chfn

# Install tools
RUN apt-get update -q 2 && DEBIAN_FRONTEND=noninteractive \
    apt-get install -y ca-certificates > /dev/null

RUN apt-get update -q 2 && DEBIAN_FRONTEND=noninteractive \
    apt-get install --yes -q 2 golang >/dev/null

RUN rm -rf /tmp/* /var/tmp/* && \
    rm -rf /var/lib/apt/lists/* && \
    rm -f /etc/dpkg/dpkg.cfg.d/02apt-speedup
