# INSTALL
Extract the tarball, then copy the binary into the monitoring plugins directory.
For example, on Debian:

	cp check_wildflyfree /usr/lib64/nagios/plugins

# EXAMPLE
Check the host wildfly.example.com with username admin and password 's3cr3t'.
Exit with critical status if 80% of memory is used.
Exit with warning if greater than 70% is used.

	check_wildflyfree -h wildfly.example.com -u admin -p s3cr3t -w 70 -c 80

See `icinga2.conf` for example Icinga2 configuration using check_wildflyfree.

# SEE ALSO
Most importantly: the man page!

	nroff -mdoc check_wildflyfree.1 | less
