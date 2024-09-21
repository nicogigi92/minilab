# minilab

The purpose of this repository is to track the evolution and architecture of my home lab. Its aim is to enable POCs to be carried out on technologies based around Kubernetes. 

I wanted to have a k8s cluster at home so I could play and test things. I wanted it to be highly available in terms of both applications and hardware. That's why I chose to buy 4 small workstations rather than a large virtualisation server. I imagine that this also allows me to benefit from good performance, but that remains to be proven.

## Hardware

- x4 Dell Optiplex 3060 MFF workstation
- x1 8 Ports Ethernet Switch  NETGEAR GS108 
- x1 Raspberry Pi 5
- x1 Rasperry Pi 3
- x7 BrandRex 6A ethernet cable

## Architecture

My lab is not exposed to the internet via my home router. This avoids dependency on the home router and reduces the amount of network configuration.
To access it, I just need to connect my PC to my ZeroTier network. As the loadbalancers are also in the network, I have access to the API server (to administer my cluster) and the ingress (to consume its services).

![Alt text](./architecture.png "Architecture")

