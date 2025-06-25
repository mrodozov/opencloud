---
status: proposed
date: 2025-06-25
author: Pascal Bleser <p.bleser@opencloud.eu>
decision-makers:
consulted:
informed:
title: "API for the Groupware Web UI"
template: https://raw.githubusercontent.com/adr/madr/refs/tags/4.0.0/template/adr-template.md
---

* Status: draft

## Context

We need a comprehensive HTTP API for the OpenCloud Web UI to provide access to the following (upcoming) modules and Groupware functionalities:

* Mail
* Contacts
* Calendar
* Tasks
* Chat
* Configuration

Current communication is done via [LibreGraph API](https://github.com/opencloud-eu/libre-graph-api) based on [Microsoft Graph](https://developer.microsoft.com/en-us/graph). An alternative is an independent API tailored to fit the specific needs of the Groupware Modules. This ADR will provide insight in the decision making process to find the most suitable and feasible solution.

## Considered Options

* LibreGraph
* JMAP
* a custom REST API

## Decision Outcome

TODO

### Consequences

TODO

### Confirmation

TODO

## Pros and Cons of the Options

### LibreGraph

* good: is already in use as the API for OpenCloud Drive operations
* neutral: does not have to follow the Microsoft Graph API, can be customized to our own needs, but in which case it becomes doubtful that there is any benefit in mimicking the Graph API in the first place if we diverge from it
* bad: not an easy API to implement, although we have libraries that take care of some of the more complex parts, such as parsing [OData](https://www.odata.org/) expressions
* bad: not tailored to our needs, and we will most probably have a lot of cases in which we have to twist the Graph API to express what the UI needs by using complex filters, which then require complex parsing in the backend in order to translate them into JMAP, as opposed to directly using an expressive and maximally matching API in the first place
* neutral: there is no compatibility benefit, since the only MUA that uses the Microsoft Graph API is Microsoft Outlook, and it is not a goal to support Microsoft Outlook as a MUA beyond standard IMAP/SMTP/CalDAV/CardDAV services (and that would be Microsoft Graph, not LibreGraph nor any customizations we would require)
* neutral: in a similar vein, we will not implement all the aspects of the Microsoft Graph API, which renders any compatibility aims moot while conserving all of the complexity drawbacks
* bad: does not support multiple accounts per user

TODO: more advantages of using LibreGraph {<j.dreyer@opencloud.eu>, <m.bartz@opencloud.eu> ?}

### JMAP

* good: very flexible protocol that can easily be implemented by clients
* good: does not require implementation efforts on the backend side
* potentially bad: potentially too flexible for its own good, as it makes it difficult to reverse-engineer the high-level meaning of a set of JMAP requests in order to capture its semantics, e.g. to implement caching or reverse indexes for performance
* neutral: JMAP will not cover 100% of the Web UI API needs (e.g. at the very least for things like configuration settings)
* good: would obviously support the full potential of JMAP and Stalwart

### Custom REST API

* good: completely tailored to the needs of the OpenCloud UI
* good: a higher-level API allows for easily understanding the semantic of each operation, which enables the potential for keeping track of data in order to implement reverse indexes and caching, if necessary to achieve functional or performance goals, as opposed to using a lower-level API such as JMAP which is maximally flexible and difficult to reverse-engineer the meaning of the operation and data
* potentially bad: does not follow any standard (besides REST), although the purpose is solely to build an API for the OpenCloud UI, not an API that is meant to be consumed by many different clients
* good: can also be tailored to the capabilities of JMAP without exposing all of its flexibility
* good: provides the potential for expanding upon what JMAP provides
* good: would support the full potential of JMAP and Stalwart since the API would be designed accordingly
