package template

var MAIN = `package ${package};

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
${imports}

/**
 * Hello world!
 *
 */
${annotations}
@SpringBootApplication
public class ${name}
{
    public static void main( String[] args )
    {
        SpringApplication.run(${name}.class, args);
    }
}
`
