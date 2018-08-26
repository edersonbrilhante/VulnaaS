# VulnaaS API

API used by VulnaaS.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

* Golang: `brew install golang` 

## Installing

#### Setting up environment variables (use your own configuration):

```
echo 'export API_HOST="192.168.50.1"' >> .env
echo 'export API_PORT="9999"' >> .env
```

```
source .env
```

## Running 

`go run server.go`

Set `:shell, path:` into your Vagrantfile to point to VulnaaS-API, as shown on the example bellow:

```
config.vm.define "vm-example" do |vm1|
    config.vm.network :private_network, ip: "192.168.50.10"
    vm1.vm.provision :shell, path: "http:localhost:9999/custom-configs/local/1001.sh", privileged: true
  end
```

## Architecture draft
