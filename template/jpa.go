package template

var JPA = `package ${package}.model;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;


@Repository
public interface ${name} extends JpaRepository<${entityFileName},String> {

}
`
