
# Cloud Foundry Service Broker for Redis Labs Enterprise Cluster(RLEC)

## Configuring the environment
### Getting the code
Get the latest version of the code by cloning the repository in the following way:
```
git clone https://github.com/RedisLabs/cf-redislabs-broker.git
cd cf-redislabs-broker
```

After you cloned the repository, please make sure that the following prerequisites are met:

* Install Go 1.5
* [Ensure your GOPATH is set correctly](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable)
* In managing dependencies, we rely on Go 1.5 Vendor Experiment. Therefore, set up a `GO15VENDOREXPERIMENT` variable to equal `1`. You can use `./bin/go` to have it set up for you.

### Building the project
To build the service broker simply run the following command:
```
./bin/build
```
After the build is completed you can locate the resulting binary in `out/redislabs-service-broker` .

### Building a Docker Image
To build a [Docker](https://www.docker.com/) image, run the following `make` command:

```console
make docker-build
```

You can also push the same Docker image with:

```console
make docker-push
```

The `Makefile` is configured by the following environment variables:

- `DOCKER_REPOY` - the registry to push to. For example, your own internal registry, [Quay](https://quay.io), or [Google Container Registry](https://cloud.google.com/container-registry/). Must have a trailing slash. Leave empty if you'd like to push to [Docker Hub](https://hub.docker.com/). Defaults to `quay.io/`
- `DOCKER_ORG` - the prefix that goes before the image name. For example, in `quay.io/my/image:tag`, `my` is the `DOCKER_ORG`. Defaults to `redislabs`
- `DOCKER_VERSION` - the tag of the image. For example, in `quay.io/my/image:tag`, `tag` is the `DOCKER_VERSION`. Defaults to the evaluation of `git-$(shell git rev-parse --short HEAD)`

### Running unit tests
It is highly advisable to execute the unit tests after the build was done.
To do so simply execute the following command:
```
./bin/test
```

### How to add a new dependency
If you would like to add a new dependency to the service broker, you can do so in the following way:

* Install [godep](https://github.com/tools/godep)
* Install the dependency (eg via `go get`) and ensure everything works fine
* `godep save ./...`
* Check that the output of `git diff vendor/ Godeps/` looks reasonable
* Commit `vendor/` and `Godeps/`


## Running the service

Start the service by running `redislabs-service-broker` pointing it to a config file in the following way:
```
redislabs-service-broker -c /path/to/config.yml
```

You can find a template of the config file under the `examples` [folder](https://github.com/RedisLabs/cf-redislabs-broker/tree/master/examples/config.yml).
This template is distributed with every release as `config.yml.template`.

You'll also need to define the following environment variables when running the broker:

- `REDIS_LABS_USERNAME` - your RedisLabs username
- `REDIS_LABS_PASSWORD` - your RedisLabs password
- `REDIS_LABS_CLUSTER_ADDR` - your RedisLabs cluster address
- `BROKER_USERNAME` - the username for the broker to use to authenticate clients (using HTTP basic auth)
- `BROKER_PASSWORD` - the password for the broker to use to authenticate clients (using HTTP basic auth)


## Using the service
To better understand how CF service brokers works please consult the the [CF documentation](http://docs.cloudfoundry.org/services/managing-service-brokers.html) .

* You can add additional configuration parameters on provisioning or updating a service, using the `-c` switch, as follows:
```
cf create-service ... -c '{"name":"myredis-db", "replication":true, "memory_size":104857600}'
```

See the RLEC API docs for the applicable parameters.

* Note that the broker is working synchronously- please wait for requests to complete.

## Logs

The service broker logs DEBUG-level info to `stdout` and errors to `stderr`.

## Internal state

The broker stores its state in a JSON file located in a `$HOME/.redislabs-broker` folder.
**NOTE:** Do not change the contents of this folder manually.

The persistence is implemented as a pluggable backend. Therefore, an option of storing the state in a database may be added in the future.
