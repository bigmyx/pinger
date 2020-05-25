[![travis](https://travis-ci.com/bigmyx/pinger.svg?branch=master)](https://travis-ci.com/github/bigmyx/pinger)


# pinger

## Simple Port Scanner

Scan a single host:

```
./pinger 127.0.0.1            
 Report for host 127.0.0.1 
 port 110       open 
 port 585       open 
 port 995       open 
 port 993       open 
 port 143       open 
 port 80        open 
 port 4767      open 
 Scanned 1 hosts in 2.098227587s
```

Scan IP block:

```
./pinger 10.0.0.0/28
 Report for host 10.0.0.11 
 port 80        open 
...
 port 8080      open 
 port 8090      open 
 Report for host 10.0.0.1 
 port 80        open 
...
 port 8080      open 
 Report for host 10.0.0.12 
 port 80        open 
... 
 port 9119      open 
 Report for host 10.0.0.10 
 port 80        open 
...
 port 3128      open 
 Scanned 4 hosts in 19.309185694s
```
