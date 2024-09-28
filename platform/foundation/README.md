# Foundation

In this folder I keep my Helm Chart configuration for cluster foundations.

I use the native helm controller of k3s to manage my charts in an elegant way. For the moment this stack is very basic but should be the big work of this project. We have:

- Longhorn to use workstation SSDs as persistent and distributed storage.
- Kube prometheus stack to have a base on which to add future monitoring. 
- An overconfiguration of traefik to make its dashboard accessible.

This project does not contain a dedicated DNS for the moment and ingress can be accessed by modifying the `/etc/hosts` file.
