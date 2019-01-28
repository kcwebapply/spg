# spg

[![GoDoc](https://godoc.org/github.com/kcwebapply/spg?status.svg)](https://godoc.org/github.com/kcwebapply/spg)
![Go Report Card](https://goreportcard.com/badge/github.com/kcwebapply/spg)
[![Release](https://img.shields.io/github/release/kcwebapply/spg.svg?style=flat-square)](https://github.com/kcwebapply/spg/release)

**`SPG`** is simple terminal tool for generating **_SpringBoot package_**  and _**Classes**_ easily and quickly.

- [Usage](#Usage)
- [Generate Package](#Generate)
- [toml Setting file](#toml)
- [generated oackage constitution](#const)

## Usage


<h2 id='Generate'>Generate Package</h2>
you can generate `SpringBoot` package with below command.
```terminal
$ spg file test.toml
Generating package spring-boot-generator !
```


<h2 id='toml'>toml Setting file</h2>
you should touch `.toml` file to select what kinds of files(`Java Class`) you want to generate automatically.

Here is the example of generating _DB-related_ file.

```toml
[App]
  name="spring-boot-generator" 
  groupId="com.kcwebapply"    
  artifactId="spring-sample"  

[Db]
  jdbc="jdbc:postgresql://localhost:5432/test" 
  driver="org.postgresql.Driver"               
  table="Purchase"                            
```

Then, after we run `spg` command, we generate package anc classes .

<h2 id='const'>Generated package constitution</h2>
In this case, package constitution is like this.

```
|--pom.xml                                       // generated automatically by default (some dependency added automatically ).
|--src
|  |--main
|  |  |--java
|  |  |  |--com
|  |  |  |  |--kcwebapply
|  |  |  |  |  |--springsample
|  |  |  |  |  |  |--SpringBootGenerator.java    // generated automatically by default.
|  |  |  |--model
|  |  |  |  |--PurchaseEntity.java               // entity generated  by Db setting (refer to Db.Table key on toml file). 
|  |  |  |  |--PurchaseRepository.java           // repository generated  by Db setting (refer to Db.Table key on toml file).
|  |  |--resources
|  |  |  |--application.properties               // generated automatically by default.
|  |--test
|  |  |--java
|  |  |  |--com
|  |  |  |  |--kcwebapply
|  |  |  |  |  |--springsample
|  |  |  |  |  |  |--SpringBootGeneratorTests.java // generated automatically by default.

```
As you see, this command generate the basic constitution of _SpringBoot_ package.

The most remarkable feature is that this library generate `Entity` class and `Repository` interface.

- `Entity` class.
```java
@Entity
@Table(name="Purchase")
public class PurchaseEntity {

}
```

- `JpaRepository` interface.
``` java
@Repository
public interface PurchaseRepository extends JpaRepository<PurchaseEntity,String> {

}
```

This command generate these _boilerplate class_  about foundational function on _`SpringBoot`_ .

#### application.properties
Depending on what you write on `.toml` file, some property is also written on `application.properties` automatically.
`application.properties` is also modified to adapt you setting.

```
spring.datasource.url=jdbc:postgresql://localhost:5432/test
spring.datasource.driver-class-name=org.postgresql.Driver
```

#### pom.xml
You add database setting on your `.toml` file, so then some dependency is added to `pom.xml` automatically.

```xml
<dependency>
     <groupId>org.springframework.boot</groupId>
     <artifactId>spring-boot-starter-data-jpa</artifactId>
</dependency>
```

Like this, when you write some library setting on `.toml` file,
dependency package is also added automatically.

<h2 id='supported'>supported function</h2>
- SpringDataJpa (Mysql, postgres)

## Install

### On macOS

```
brew tap kcwebapply/spg
brew install spg
```



