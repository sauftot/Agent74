# Agent74
Sipmle reverse agent proxy for tcp connections.

# How it works
Used to forward a tcp service sat behind a firewall to external clients. This is done by caching client connections/packets on an internet facing, exposed small proxy server (cheap VPS) running Agent74server which is then polled by the Agent74client in order to relay the packets to the tcp service via localhost.
