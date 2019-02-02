package template

var ENTITY = `package ${package}.model;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;


@Entity
@Table(name=${tableName})
public class ${name} {

  @Id
  private String id;

  public String getId(){
      return id;
  }

  public void setId(String id){
      this.id = id;
  }

}
`
