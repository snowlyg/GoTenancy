# QOR

English Chat Room: [![Join the chat at https://gitter.im/qor/qor](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/qor/qor?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

中文聊天室： [![加入中国Qor聊天室 https://gitter.im/qor/qor/china](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/qor/qor/china)

[![Build Status](https://semaphoreci.com/api/v1/theplant/qor/branches/master/badge.svg)](https://semaphoreci.com/theplant/qor)

**For security issues, please send us an email to security@getqor.com and give us time to respond BEFORE posting as an issue or reporting on public forums.**

## What is QOR?

QOR is a set of libraries written in Go that abstracts common features needed for business applications, CMSs, and E-commerce systems.

This is a complete rewrite in Go, of the original QOR, which was a proprietary framework written in Ruby on Rails and used internally at [The Plant](https://theplant.jp). QOR 1.0 is the first version to be open sourced and distributed under the MIT license.

### What QOR is not

QOR is not a "boxed turnkey solution". You need proper coding skills to use it. It's designed to make the lives of developers easier when building complex EC systems, not providing you one out of the box.

## Documentation

<https://doc.getqor.com/>


## The modules

* [Admin](https://GoTenancy/libs/admin) - The core part of QOR system, will generate an admin interface and RESTFul API for you to manage your data

* [Publish](https://GoTenancy/libs/publish) - Providing a staging environment for all content changes to be reviewed before being published to the live system

* [Transition](https://GoTenancy/libs/transition) - A configurable State Machine: define states, events (eg. pay order), and validation constraints for state transitions

* [Media Library](https://GoTenancy/libs/media_library) - Asset Management with support for several cloud storage backends and publishing via a CDN

* [Worker](https://GoTenancy/libs/worker) (Batch processing) - A process scheduler

* [Exchange](https://GoTenancy/libs/exchange) - Data exchange with other business applications using CSV or Excel data

* [Internationalization](https://GoTenancy/libs/i18n) (i18n) - Managing and (inline) editing of translations

* [Localization](https://GoTenancy/libs/l10n) (l10n) - Manage DB-backed models on per-locale basis, with support for defining/editing localizable attributes, and locale-based querying

* [Roles](https://GoTenancy/libs/roles) - Access Control

* And more [https://GoTenancy/libs](https://GoTenancy/libs)

## Live DEMO

* Live Demo [http://demo.getqor.com/admin](http://demo.getqor.com/admin)
* Source Code of Live Demo [https://GoTenancy/libs/qor-example](https://GoTenancy/libs/qor-example)

## Frontend Development

Requires [Node.js](https://nodejs.org/) and [Gulp](http://gulpjs.com/) for building frontend files

```bash
npm install && npm install -g gulp
```

- Watch SCSS/JavaScript changes: `gulp`
- Build Release files: `gulp release`

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
