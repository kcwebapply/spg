# spg

[![GoDoc](https://godoc.org/github.com/kcwebapply/spg?status.svg)](https://godoc.org/github.com/kcwebapply/spg)
![Go Report Card](https://goreportcard.com/badge/github.com/kcwebapply/spg)
[![Release](https://img.shields.io/github/release/kcwebapply/spg.svg?style=flat-square)](https://github.com/kcwebapply/spg/release)

**`SPG`** is simple terminal tool for generating **_SpringBoot package_**  and _**Classes**_ easily and quickly.

- [Usage](#Usage)
- [Generate Package](#Generate Package)
- [toml Setting file](#toml Setting file)

## Usage

### Generate Package
you can generate `SpringBoot` package with below command.
```terminal
$ spg file test.toml
Generating package spring-boot-generator !
```



### toml Setting file
you should touch `.toml` file to select what kinds of files(`Java Class`) you want to generate automatically.

Here is the example of generating _DB-related_ file.

```toml
[App]
  name="spring-boot-generator"

[Db]
  jdbc="jdbc:postgresql://localhost:5432/test"
  driver="org.postgresql.Driver"
  table="Purchase"
```

Then, after we run `spg` command, we generate package anc classes .

### Generated package constitution
In this case, package constitution is like this.

```
/spring-boot-generator
|--pom.xml                              // generated automatically by default (some dependency added automatically ).
|--src
|  |--main
|  |  |--java
|  |  |  |--SpringBootGenerator.java    // generated automatically by default.
|  |  |  |--model
|  |  |  |  |--PurchaseEntity.java      // entity generated  by Db setting (refer to Db.Table key on toml file). 
|  |  |  |  |--PurchaseRepository.java  // repository generated  by Db setting (refer to Db.Table key on toml file).
|  |  |--resources
|  |  |  |--application.properties      // generated automatically by default.
```
As you see, this library generate the basic constitution of _SpringBoot_ package.

##### application.properties
Depending on what you write on `.toml` file, some property is also written on `application.properties` automatically.
`application.properties` is also modified to adapt you setting.

```
spring.datasource.url=jdbc:postgresql://localhost:5432/test
spring.datasource.driver-class-name=org.postgresql.Driver
```

##### pom.xml
if you add database setting on your `.toml` file, then some dependency is added to `pom.xml` automatically.

```xml
<dependency>
     <groupId>org.springframework.boot</groupId>
     <artifactId>spring-boot-starter-data-jpa</artifactId>
</dependency>
```



## Install

### On macOS

```
brew tap kcwebapply/spg
brew install spg
```



