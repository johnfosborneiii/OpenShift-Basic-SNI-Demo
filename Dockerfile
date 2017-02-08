FROM openshift/origin-base
ADD server /usr/sbin/server
CMD ["/usr/sbin/server"]
