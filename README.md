# CloudNativePG Interface (CNPG-I)

**Status:** Experimental

The **CloudNativePG Interface (CNPG-I)** introduces a modular and extensible
approach to integrating plugins with the
[CloudNativePG operator for PostgreSQL](https://github.com/cloudnative-pg/cloudnative-pg),
enabling greater flexibility and ease of feature development in the
CloudNativePG ecosystem.

## What is CNPG-I?

The **CloudNativePG Interface** is a **gRPC-based protocol** that defines a
standardized interface between the CloudNativePG operator and external plugins.
This approach empowers developers and organizations to extend the operator's
functionality without the need to fork its codebase.

For a detailed protocol specification, refer to the [protocol.md](docs/protocol.md) file.

## Why CNPG-I?

Over the years, CloudNativePG's codebase has grown significantly, creating
challenges for developers contributing to the project and increasing the time
required to deliver new features.

To address this, CNPG-I provides a **modular plugin architecture** inspired by
the success of Kubernetes' [Container Storage Interface (CSI)](https://github.com/container-storage-interface/spec/blob/master/spec.md).

CNPG-I fosters a thriving ecosystem by making it easier to integrate new
capabilities through independent plugins, allowing for faster innovation and
simplified maintenance.

### Key Benefits

1. **Reduced Complexity**: Delegate responsibilities to external plugins,
   reducing the cognitive load and size of the CloudNativePG core codebase.
2. **Faster Feature Delivery**: Accelerate development by fostering independent
   plugin creation, prototyping, and ecosystem growth.
3. **Standardized gRPC API**:
   Provide a robust, gRPC-based API specification for customizing operator and
   operand behavior, complete with compliance testing tools.
4. **Ecosystem Expansion**: Encourage innovation by enabling developers to
   create enhancements as separate projects, avoiding forks and lowering barriers
   to entry.
5. **Ease of Adoption**: Simplify the move to CloudNativePG for Postgres and
   Kubernetes users by facilitating custom enhancements and integrations.

## Use Cases for Plugins

CNPG-I enables the creation of plugins for diverse use cases, such as:

- **WAL Management**
- **Backup and Recovery**
- **Logging and Auditing**
- **Metrics Export**
- **Authentication and Authorization**
- **Extension Management**
- **Instance Lifecycle Management**
- **Configuration Management**

By leveraging plugins, developers can extend CloudNativePG’s functionality
without modifying its core code.

## Projects Built with CNPG-I

Explore real-world applications of CNPG-I:

- [CNPG-I Hello World](https://github.com/cloudnative-pg/cnpg-i-hello-world/):
  A simple example plugin demonstrating CNPG-I.
- [Barman Cloud CNPG-I Plugin](https://github.com/cloudnative-pg/plugin-barman-cloud):
  A plugin for seamless integration with Barman Cloud for Continuous Backup and
  Recovery using object stores.

---

## Trademarks

*[Postgres, PostgreSQL and the Slonik Logo](https://www.postgresql.org/about/policies/trademarks/)
are trademarks or registered trademarks of the PostgreSQL Community Association
of Canada, and used with their permission.*
