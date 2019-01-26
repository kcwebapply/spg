# spg

[![GoDoc](https://godoc.org/github.com/kcwebapply/spg?status.svg)](https://godoc.org/github.com/kcwebapply/spg)
![Go Report Card](https://goreportcard.com/badge/github.com/kcwebapply/spg)
[![Release](https://img.shields.io/github/release/kcwebapply/spg.svg?style=flat-square)](https://github.com/kcwebapply/spg/release)

**`SPG`** is simple terminal tool for generating **_SpringBoot package_**  and _**Classes**_ easily and quickly.

## Usage

#### Generate Package
you can generate `SpringBoot` package with below command.
```terminal
$ spg file test.toml
Generating package spring-boot-generator !
```



#### toml Setting file
you should touch `.toml` file to select what kinds of files you want to generate automatically.

Here is the example of generating _DB-related_ file.
```toml
[App]
  name="spring-boot-generator"

[Db]
  jdbc="jdbc:postgresql://localhost:5432/test"
  driver="org.postgresql.Driver"
  table="Purchase"
```

In this case, package components this.

```
/spring-boot-generator
|--pom.xml
|--src
|  |--main
|  |  |--java
|  |  |  |--SpringBootGenerator.java    // generated automatically by default.
|  |  |  |--model
|  |  |  |  |--PurchaseEntity.java      // entity generated automatically by Db setting (refer to `table` key). 
|  |  |  |  |--PurchaseRepository.java  // repository generrated automatically by Db setting (refer to `table` key).
|  |  |--resources
|  |  |  |--application.properties      // generated automatically by default.
```

`application.properties` is also modified to adapt you setting.

```
spring.datasource.url=jdbc:postgresql://localhost:5432/test
spring.datasource.driver-class-name=org.postgresql.Driver
```

## Install

### On macOS

```
brew tap kcwebapply/spg
brew install spg
```



