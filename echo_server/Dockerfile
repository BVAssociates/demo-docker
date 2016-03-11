FROM bvassociates/i386-golang
MAINTAINER Vincent Bauchart <bauchart@bvassociates.fr>

ADD src/echo_server.go /tmp/

EXPOSE 12321

# Use the runit init system.
CMD ["/usr/bin/go", "run", "/tmp/echo_server.go","/etc/hostname"]
