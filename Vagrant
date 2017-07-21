# vi: set ft=ruby :
 
Vagrant.configure("2") do |config|
    config.vm.box = "ubuntu/xenial64"
    config.vm.synced_folder ENV['HOME'], "/myhome", type: "nfs"
    #config.vm.synced_folder '.', "/myfolder", type: "nfs"

    config.vm.provider "virtualbox" do |v|
      v.memory = 4096
      v.cpus = 2
    end
 
    config.vm.define "tc-rocksdb" do |node|
      node.vm.hostname = "tc-rocksdb"
      node.vm.network :private_network, ip: "192.168.33.99"

      # Use NFS for shared folders for better performance
      node.vm.synced_folder '.', '/vagrant', nfs: true
      #config.vm.synced_folder ENV['HOME'], "/myhome", type: "nfs"

      node.vm.provision :shell, inline: "sed 's/127\.0\.0\.1.*master.*/192\.168\.99\.10 master/' -i /etc/hosts"
    end
 
end
