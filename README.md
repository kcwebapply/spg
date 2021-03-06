# spg



[![GoDoc](https://godoc.org/github.com/kcwebapply/spg?status.svg)](https://godoc.org/github.com/kcwebapply/spg)
![Go Report Card](https://goreportcard.com/badge/github.com/kcwebapply/spg)
[![Release](https://img.shields.io/github/release/kcwebapply/spg.svg?style=flat-square)](https://github.com/kcwebapply/spg/release)


<img  align="right" src="image/spg.png" width="200px">

**`SPG`** is simple terminal tool for generating **_SpringBoot package_**  and _**Classes**_ easily and quickly.


<img src="image/sample.gif" width="500" height="300">.


- [Why Spg?](#Spg)
- [Usage](#Usage)
  - [Generate Package](#Generate)
  - [toml Setting file](#toml)
  - [Generated oackage constitution](#const)
  - [Generate toml file](#toml)
- [Supported function](#supported)
- [Installation](#install)
  - [mac Os App](#mac)

<h1 id="Spg">Why Spg?</h1>

If you want to generate _SpringBootPackage_  , you can take 2 ways

- `mvn -B archetype:generate` command.
- using [spring-initializr](https://start.spring.io/)

<h2>mvn commands</h2>

But, `mvn` commands needs lots of typing and _generated package_ is only **_JavaPackage_**, it is not **_SpringBootPackage_** .

<h2>Spring-initializr</h2>

`Spring-initializr` is a good way of generating **_SpringBootPackage_**.

But `Spring-initializr` can be only used on online.

Also,what `Spring-initializr` do about `dependency` is only adding `dependency` tag on `pom.xml` in many case.
So, If you add `dependency` on `Spring-initializr` and download it, its package may imcomplete one .

<h2>Then, What I want is this. Spg is this.</h2>

- Cli which only takes `few typing`
- Generate _SpringBootPackage_ .
- Can be used offline.
- Generate completed application depends on what function (ex, API, DB, Scheduler) we want to use .

To do that, I make `spg` commands.


<h1 id="Usage">Usage</h1>


<h2 id='Generate'>Generate Package</h2>

you can generate `SpringBoot` package with below command very quickly.

```terminal
spg gen -g "com.sample" -a "sampleproject" -name "sample-project"
```

then, you can generate _SpringBootPackage_ like this.

```
├── pom.xml
└── src
    ├── main
    │   ├── java
    │   │   └── com
    │   │       └── sample
    │   │           └── sampleproject
    │   │               └── SampleProject.java
    │   └── resources
    │       └── application.properties
    └── test
        └── java
            └── com
                └── sample
                    └── sampleproject
                        └── SampleProjectTests.java
```


This is the basic components of _SpringBootPackage_ .

you can also generate _SpringBootPackage_ from `.toml` file.

```terminal
$ spg genf test.toml
Generating package spring-boot-generator !
```

`.toml` file formats are below.


<h2 id='toml'>toml Setting file</h2>



you should touch `.toml` file to select what kinds of  _Java-Class_ files you want to generate automatically.

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

<h2 id='toml'>Generating toml file</h2>

you can generate `.toml` file with `spg touch` command.

```terminal
$ spg touch -artifactId spring-boot-generator -g com.test -n spring-boot-generator
Generating spg.toml file completed!
```
generate file like this.
```toml
// spg.toml
[App]
  name="spring-boot-generator"
  groupId="com.test"
  artifactId="spring-boot-generator"
  springVersion="2.1.1.RELEASE"
  javaVersion="1.8"
```

<h1 id='supported'>supported function</h2>

Here is the list of supported function and `.toml` setting.
- [SpringData Jpa](#db)
- [SpringScheduker](#db)

<h3 id='db'>SpringDataJpa (Mysql, postgres)</h3>

```toml
[Db]
  jdbc="jdbc:postgresql://localhost:5432/test"
  driver="org.postgresql.Driver"               
  table="Purchase"         

```

<h3 id='task'>SpringScheduler</h3>

```toml
[task]
  schedule = "0 * * * * *"
  zone="Asia/Tokyo"
```

<h1 id='install'>Install</h1>

<h2 id='mac'>On macOS</h2>

```
brew tap kcwebapply/spg
brew install spg
```
