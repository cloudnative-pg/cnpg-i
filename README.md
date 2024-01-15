# CloudNativePG Interface (CNPG-I)

**Status:** PROPOSAL

## Why do we need CNPG-I?

The CloudNativePG code base has sensibly grown over the years, increasing the
cognitive load required to work with it and creating a bottleneck in the
submission of new pull requests in the open source project itself.

The risk we are currently facing is that the delivery of new features is slowed
down further and further as the project grows both in code base and adoption.

As with every successful open source project indeed, everyday maintainers need
to face the dilemma to reject or approve a pull request in the code base
knowing that, when it gets committed, it’s forever.

We need to find a different and more pluggable way of adding new behaviors and
capabilities to the operator, not necessarily in the open source space.

The proposal here is to follow a pattern that’s already been successfully
implemented in Kubernetes through the
[Container Storage Interface (CSI)](https://github.com/container-storage-interface/spec/blob/master/spec.md),
and give birth to a new open source initiative within CloudNativePG called the
“CloudNativePG Interface (CNPG-I)”.

The CNPG-I standard API, based on gRPC, should enable new organizations to join
EDB (original creator of CloudNativePG) and develop a plugin or a collection of
them - instead of creating, maintaining, and distributing a fork of
CloudNativePG.

Some of the advantages of CNPG-I can bring are:

* Remove complexity and responsibility from CloudNativePG open source code
  base, by delegating them to external plugins that can run as sidecar
  containers alongside Postgres
* Remove bottleneck from CloudNativePG core team in terms of delivery of new
  features, increasing velocity, fostering prototyping, and nurturing an
  independent ecosystem of plugins
* A standard gRPC-based API specification covering the customizable aspects of
  the operator and the operands, and providing a way to test the compliance
  with it
* Encourage dynamic growth of the ecosystem through customizations that can be
  independently developed within the operands, as separate projects instead of
  operator forks, with a lower required cognitive load
* Facilitate the move to CloudNativePG for other Postgres and Kubernetes
  companies.

## Some examples of plugins

For example, plugins might involve:

* WAL management (archive and restore)
* Backup and recovery management
* Logging and auditing
* Export of metrics
* Authentication and authorizations
* Management of extensions
* Management of an instance lifecycle
* Management of the configuration of an instance

## Trademarks

*[Postgres, PostgreSQL and the Slonik Logo](https://www.postgresql.org/about/policies/trademarks/)
are trademarks or registered trademarks of the PostgreSQL Community Association
of Canada, and used with their permission.*
