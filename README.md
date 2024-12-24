# CloudNativePG Interface (CNPG-I)

**Status:** Experimental

## Protocol

The CloudNativePG Interface (CNPG-I) is a gRPC-based protocol that defines the
interface between the CloudNativePG operator and its plugins.

For the full protocol specification, refer to the [protocol.md](docs/protocol.md) file.

## Rationale

The CloudNativePG code base has grown significantly over the years, increasing
the cognitive load required to work with it and creating challenges in the
submission of new pull requests in the open source project.

As the project grows both in code base and adoption, it is important to ensure
that the delivery of new features remains efficient.

To address this, we propose a more modular approach to adding new behaviors and
capabilities to the operator, inspired by the successful implementation of the
[Container Storage Interface (CSI)](https://github.com/container-storage-interface/spec/blob/master/spec.md)
in Kubernetes.

The CloudNativePG Interface (CNPG-I) standard API, based on gRPC, aims to enable
developers and organizations to create plugins or collections of them, rather
than creating, maintaining, and distributing a fork of CloudNativePG.

Some of the benefits of CNPG-I include:

* Reducing complexity and responsibility within the CloudNativePG open source
  code base by delegating them to external plugins that can run as sidecar
  containers alongside Postgres
* Increasing the velocity of feature delivery by removing bottlenecks from the
  CloudNativePG core team, fostering prototyping, and nurturing an independent
  ecosystem of plugins
* Providing a standard gRPC-based API specification that covers customizable
  aspects of the operator and the operands, and offering a way to test
  compliance with it
* Encouraging dynamic growth of the ecosystem through customizations that can be
  independently developed within the operands, as separate projects instead of
  operator forks, with a lower required cognitive load
* Facilitating the move to CloudNativePG for other Postgres and Kubernetes
  companies.

### Use cases for plugins

For example, plugins might involve:

* WAL management
* Backup and recovery management
* Logging and auditing
* Export of metrics
* Authentication and authorization
* Management of extensions
* Management of an instance lifecycle
* Management of the configuration of an instance

## Projects using CNPG-I

* [CNPG-I Hello World](https://github.com/cloudnative-pg/cnpg-i-hello-world/)
* [Barman Cloud CNPG-I Plugin](https://github.com/cloudnative-pg/plugin-barman-cloud)

## Trademarks

*[Postgres, PostgreSQL and the Slonik Logo](https://www.postgresql.org/about/policies/trademarks/)
are trademarks or registered trademarks of the PostgreSQL Community Association
of Canada, and used with their permission.*
