# tc-ubuntu-xenial64-rocksdb

Ubuntu 14.10 Vagrant Development machine for RocksDB development

## Getting started

```bash
git clone https://github.com/topconnector/tc-ubuntu-xenial64-rocksdb.git
cd tc-ubuntu-xenial64-rocksdb
cd single-machine
```

You must have the following installed:

* Virtualbox >= 5.1.22

  Download and install from https://www.virtualbox.org/.

* Vagrant >= 1.9.7

  Download and install from https://www.vagrantup.com/.

* vagrant-vbguest Vagrant plugin
  automatically installs the host's VirtualBox Guest Additions on the guest system.

  Install by running: 

```bash
   vagrant plugin install vagrant-vbguest
   vagrant vbguest --do install
```

 
* update Vagrant box

  Install by running: 
    
```bash
    vagrant box update
```
    
* run Virtual machine (VM)

  Install by running: 
  
```bash
    vagrant up
```

