package com.xie.java;

import org.junit.jupiter.api.Test;

import java.util.Collections;
import java.util.TreeMap;

/**
 * @author xie4ever
 * @date 2021/5/12 17:38
 */
public class Sort {

    @Test
    public void test() {
        TreeMap<Integer, String> pairs = new TreeMap<>(
                Collections.reverseOrder()
        );

        pairs.put(2, "B");
        pairs.put(1, "A");
        pairs.put(3, "C");

        System.out.println(pairs);
    }
}
