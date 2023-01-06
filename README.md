# ORM_test

Compare ruby on rails and golang db record updating loop with orm.

## Ruby on rails spec
- ruby: 3.2.0
- rails(ORM): ~> 7.0.3 (7.0.4)

## Go spec
- go: 1.19.4
- ent(ORM): 0.11.4

## Database
- posgresql: 13.7

## Test db scheme
```
CREATE TABLE data
id int PRIMARY_KEY,
contents text,
created_at datetime,
updated_at datetime
```
