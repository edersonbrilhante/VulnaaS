# VulnaaS API

API used to will make security tests inside a CI.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

* Golang: `brew insall golang` 

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

## Architecture draft

## MongoDB draft

