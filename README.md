A plugin for Nagios, Icinga et al. to monitor a Wildfly server's JVM heap usage.The check status is set based on the percentage of heap used. If outside threshold values, the status is set to warning or critical accordingly.

Here is some example output, including performance metrics for pretty graphs:

	5% of heap memory used (12345678 of 4567890 bytes) | heap_used_percent=5, heap_used_bytes=12345678

# INSTALL
Download the latest version from [Icinga Exchange](ttps://exchange.icinga.com/otl/check_wildflyfree/releases)

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

The mini library developed for this plugin: [git.sr.ht/~otl/wildfly](https://git.sr.ht/~otl/wildfly)
