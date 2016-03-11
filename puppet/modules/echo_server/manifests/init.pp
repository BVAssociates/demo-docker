# == Class: echo_server
#
# Install binary echo_server and its config file
#
# === Parameters
#
# Document parameters here.
#
# [*echo_message*]
#   The message set to be echoed
#
# === Examples
#
#  class { 'echo_server':
#    echo_message => "Puppet was here"
#  }
#
# === Authors
#
# Vincent Bauchart <bauchart@bvassociates.fr>
#
# === Copyright
#
# Copyright 2016 BV Associates
#
class echo_server (
    $echo_message = '',
) {

    file { '/usr/local/bin/echo_server':
        ensure => present,
        owner  => 'user1',
        group  => 'group1',
        mode   => '0755',
        source => "puppet:///modules/echo_server/bin/echo_server",
    }

    file { '/etc/echo_server.conf':
        ensure  => present,
        owner   => 'user1',
        group   => 'group1',
        mode    => '0755',
        content => template('echo_server/echo_server.conf.erb')
    }

    user { 'user1':
        ensure => present,
        gid    => ['group1'],
    }

    group { 'group1':
        ensure => present,
    }

}
