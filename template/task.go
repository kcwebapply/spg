package template

var TASK = `package ${package}.task;

import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

@Component
public class Task{

    @Scheduled(cron = ${schedule}, zone = ${zone})
    public void cronJob(){

    }

}
`
