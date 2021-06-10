package com.xie.java;

import java.util.HashMap;
import java.util.Map;

/**
 * @author xie4ever
 * @date 2021/6/1 10:08
 */
public class Test {

    class employee {
        public long getID() {
            return ID;
        }

        public void setID(long ID) {
            this.ID = ID;
        }

        public employee(long ID, String name) {
            this.ID = ID;
            this.name = name;
        }

        private long ID;
        private String name;
    }

    Map<employee, Double> map = new HashMap<>();

    @org.junit.jupiter.api.Test
    public void GetHighestSalary() {
        Map<employee, Double> map = new HashMap<>();
        map.put(new employee(1, "tom"), (double) 999);
        map.put(new employee(2, "bob"), (double) 1000);
        map.put(new employee(3, "alice"), (double) 1001);
        double salary = 0;
        for (employee employee : map.keySet()) {
            double tmp = map.get(employee);
            if (tmp > salary) {
                salary = tmp;
            }
        }
        System.out.println(salary);
    }

    @org.junit.jupiter.api.Test
    public void GetSum() {
        map.put(new employee(1, "tom"), (double) 999);
        map.put(new employee(2, "bob"), (double) 1000);
        map.put(new employee(3, "alice"), (double) 1001);
        int num = 0;
        for (employee employee : map.keySet()) {
            double tmp = map.get(employee);
            if (tmp > 1000) {
                num++;
            }
        }
        System.out.println(num);
    }
}
