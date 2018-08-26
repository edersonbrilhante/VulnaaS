# VulnaaS API

API used by VulnaaS.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

* Golang: `brew install golang` 

## Installing

#### Setting up environment variables (use your own configuration):

Don't forget to change this password!

```
echo 'export MONGO_HOST="192.168.50.5"' >> .env
echo 'export MONGO_DATABASE_NAME="vulnaasDB"' >> .env
echo 'export MONGO_DATABASE_USERNAME="vulnaas"' >> .env
echo 'export MONGO_DATABASE_PASSWORD="superENVPassword"' >> .env
```

```
source .env
```

## Running 

`go run server.go`

Set `:shell, path:` to point to VulnaaS-API, as shown on the example bellow:

```
config.vm.define "vm-example" do |vm1|
    vm1.vm.provision :shell, path: "http:localhost:9999/custom-configs/local/1001.sh", privileged: true
  end
```

## Architecture draft

## MongoDB draft

