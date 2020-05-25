[![Build Status](https://travis-ci.com/bigmyx/pinger.svg?branch=master)](https://travis-ci.com/bigmyx/pinger)
[![Codefresh build status]( https://g.codefresh.io/api/badges/pipeline/bigmyx/default%2Fpinger?key=eyJhbGciOiJIUzI1NiJ9.NWVjYjA2Y2M5NDdkYjE5ZDU1Nzk1MmU0.LuAwzjeXQFONvMSpkC98O7qM76gSgS1PmvLslScUJ44&type=cf-1)]( https%3A%2F%2Fg.codefresh.io%2Fpipelines%2Fpinger%2Fbuilds%3Ffilter%3Dtrigger%3Abuild~Build%3Bpipeline%3A5ecb075e13505734c6126241~pinger)


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
