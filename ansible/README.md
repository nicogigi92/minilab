# Ansible

The [inventory.yml](./inventory.yml) file is based on environment variables. Before launching a playbook, I fill in the .env file based on the [.env.sample](.env.sample) template. 

Then, after sourcing my .env file, I can launch my playbooks :

```bash
source .env
ansible-playbook -i inventory.yml ./sub-folder/playbook.yml
```

## Loadbalancers
The [loadbalancer](./loadbalancer) subfolder contains the configuration of my two loadbalancers.
- The playbook starts by installing and configuring HAProxy. It configures two different backends, one for the API server on port 6443 and the other for the cluster services on port 443. 

    *Nb: By default, k3s installs traefik, which exposes a loadbalancer service on all nodes' IPs.*

- It then installs Keepalived and configure it to expose a VIP in the ZeroTier network. A simple regex is used to detect the zerotier interface to use.

As Raspberry OS does not by default ask for the password to gain privileges, I can run the playbook as follows: 

```bash
source .env
ansible-playbook -i inventory.yml ./loadbalancer/playbook.yml
```
**To do :**
- [x] Configuring a backup and exhibiting a VIP in the ZeroTier network
- [ ] Integrate automatic connection to the ZeroTier network
- [ ] Containerise the two services and use docker-compose or another similar tool for deployment

## K3S

The k3s folder allows me to uninstall and install k3s on my hosts at my request.

- The playbook starts by initialising the cluster on the first master and immediately download the kubeconfig file
- Then it installs k3s in server mode on the other two
- Finally it installs k3s in agent mode on the last node
- The uninstall playbook uses the dedicated k3s script to uninstall the cluster

I create and kill my clusters with the following commands, using the -K option to specify the privilege escalation password : 

```bash
source .env
ansible-playbook -i inventory.yml ./k3s/install.yml -K
ansible-playbook -i inventory.yml ./k3s/uninstall.yml -K
```

I like this simple configuration because it allows me to distribute etcd and kube-api server on three nodes while still being able to run pods on my master nodes. For the little hardware and load I have, it's perfectly acceptable. 

*Nb: We install k3s in no-kube-proxy mode to make room for the Cilium data plane, the configuration of which is specified in the platform folder* 

**To do :**
- [x] Ensure HA and fully automated deployment
- [x] Install k3s in no-kube-proxy mode to allow Cilium data plane
- [ ] Add a minimum of security by running CIS hardening scripts, for example.
- [ ] Add an ansible script for upgrading k8s and even the OS