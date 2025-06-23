---
title: "WebUi API"
---

* Status: draft

## Context

We need a comprehensive and flexible HTTP API for the Opencloud WebUI to provide access to the following (upcoming) modules and functionalities of a so called groupware:

* Mail
* Contacts
* Calendar
* Tasks
* Chat
* Configuration

Current communication is done via [LibreGraph](https://github.com/opencloud-eu/libre-graph-api) based on [Microsoft Graph](https://developer.microsoft.com/en-us/graph). An alternative is an independent API tailored to fit the specific needs of the Groupware Modules. This ADR will provide insight in the decision making process to find the most suitable and feasible solution.

