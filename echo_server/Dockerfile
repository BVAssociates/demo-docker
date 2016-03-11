FROM bvassociates/i386-golang
MAINTAINER Vincent Bauchart <bauchart@bvassociates.fr>

# Copy GO file into container
ADD src/echo_server.go /tmp/

# Compilation
RUN /usr/bin/go build -o /tmp/echo_server "/tmp/echo_server.go"


# defaults params for execution

EXPOSE 12321
CMD ["/tmp/echo_server","/etc/hostname"]
