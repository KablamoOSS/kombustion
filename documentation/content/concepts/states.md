+++
title = "CloudFormation States"
description = "The interpretation of various CloudFormation States during task executions."
date = "2018-06-29T00:00:00+10:00"
weight = 20
draft = true
bref = "The interpretation of various CloudFormation States during task executions."
toc = true
layout = "docs"
+++

 |**Create Stack**|**Update Stack**|**Delete Stack**
:-----:|:-----:|:-----:|:-----:
CREATE\_COMPLETE|0|0|1
CREATE\_IN\_PROGRESS|~|~|~
CREATE\_FAILED|1|1|1
DELETE\_COMPLETE|1|1|0
DELETE\_FAILED|1|1|1
DELETE\_IN\_PROGRESS|~|~|~
REVIEW\_IN\_PROGRESS|1|1|1
ROLLBACK\_COMPLETE|1|1|1
ROLLBACK\_FAILED|1|1|1
ROLLBACK\_IN\_PROGRESS|~|~|~
UPDATE\_COMPLETE|0|0|1
UPDATE\_COMPLETE\_CLEANUP\_IN\_PROGRESS|~|~|~
UPDATE\_IN\_PROGRESS|~|~|~
UPDATE\_ROLLBACK\_COMPLETE|0|1|1
UPDATE\_ROLLBACK\_COMPLETE\_CLEANUP\_IN\_PROGRESS|~|~|~
UPDATE\_ROLLBACK\_FAILED|1|1|1
UPDATE\_ROLLBACK\_IN\_PROGRESS|~|~|~ 
