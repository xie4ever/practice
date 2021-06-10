package com.xie.java;

import org.junit.jupiter.api.Test;

/**
 * @author xie4ever
 * @date 2021/5/12 17:51
 */
public class Exception {

    class A extends java.lang.Exception{}
    class B extends A{}

    @Test
    public void test() {
        try {
            throw new A();
        } catch (B b) {
            System.out.println("caught B");
            return;
        } catch (java.lang.Exception e) {
            System.out.println("caught Exce");
            return;
        } finally {
            System.out.println("hello");
        }
    }
}
