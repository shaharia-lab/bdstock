FROM scratch
COPY bdstock /usr/bin/bdstock
ENTRYPOINT ["/usr/bin/bdstock"]