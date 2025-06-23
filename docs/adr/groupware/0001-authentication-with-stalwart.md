---
title: "Authentication with Stalwart"
---

* Status: draft

## Context

In a groupware environment, not every user will always use the OpenCloud UI to read their emails, some will resort to other [MUAs (Mail User Agents)](https://en.wikipedia.org/wiki/Email_client) that support a subset of features, use older protocols (IMAP, POP, SMTP, CalDAV, CardDAV) and lesser authentication methods (basic authentication). Those email clients will talk to Stalwart directly, as opposed to the OpenCloud UI which will make use of APIs of the OpenCloud Groupware service, since those protocols are provided by Stalwart and implementing them in OpenCloud would offer very little benefits, but definitely a lot of (almost completely) unnecessary effort.

Those protocols and operations that bypass the OpenCloud UI also need to be authenticated, this in and by Stalwart, and we need to find the best fitting approach that fulfills most or all of the following constraints:

### Single Provisioning

We want to avoid multiple provisioning of users, groups, passwords and other resources as much as possible.
While it is possible to have e.g. OpenCloud's user management also perform [Management API](https://stalw.art/docs/category/management-api/) calls, one still inevitably ends up in situations where users, user passwords, or other resources are not in sync, which becomes complex to debug and fix, and should thus be avoided if possible.

To do so, we should strive to have a single source of truth regarding users, their passwords, and similar resources and attributes such as groups, roles, application passwords, etc...

### Attack Detection

Coordinated attacks such as [denial of service](https://en.wikipedia.org/wiki/Denial-of-service_attack) attempts don't necessarily focus on a single protocol but are commonly multi-pronged, e.g. by brute forcing the [OIDC API](https://www.keycloak.org/docs/latest/authorization_services/index.html#token-endpoint), the OpenCloud Groupware API, IMAP and SMTP, *DAV protocols, etc...

In order to detect those as well as to quickly react by blacklisting clients that are identified to attempt such attacks, it is useful to have a single authentication service for all the components of the system, all protocols, all clients (e.g. [PowerDNS Weakforced](https://github.com/PowerDNS/weakforced), [Nauthilus](https://nauthilus.org/), ...)

Furthermore, such services typically make use of [DNSBL/RBL services](https://en.wikipedia.org/wiki/Domain_Name_System_blocklist) that allow IP addresses of botnets to be blocked across many services of many providers as a shared defense mechanism.

As a bonus, a centralized authentication component can also provide metrics and observability capabilities across all those protocols.

### Custom Authentication Implementations

Some customers might want custom authentication implementations to integrate with their environment, in which case we would want those to be done once and in the technology stack we're all most familiar with (thus as a service in Go in the OpenCloud framework, and not e.g. a Lua script in Nauthilus, or a Rust plugin in Stalwart, etc...)

### Custom Authentication Implementations

[Stalwart Mail Server](https://stalw.art/) has various options for User Authentication. However, the integration of Stalwart into the OpenCloud Groupware as a main source of Data needs an even more flexible way for authentication and provisioning of users. Future possible integration into existing environments of customers and partners should be considered.
