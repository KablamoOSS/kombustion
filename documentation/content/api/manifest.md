+++
title = "kombustion.yaml"
description = ""
date = "2018-06-14T00:00:00+10:00"
weight = 20
draft = false
bref = ""
toc = true
layout = "docs"
+++



```yaml
name: Test
region: ap-southeast-2
plugins:
  github.com/KablamoOSS/kombustion-plugin-serverless@0.1.0:
    name: github.com/KablamoOSS/kombustion-plugin-serverless
    version: 0.1.0
    alias: ""
environments:
  production:
    accountIDs:
    - "13521354"
    parameters:
      ENVIRONMENT: production
hideDefaultExports: false
```
