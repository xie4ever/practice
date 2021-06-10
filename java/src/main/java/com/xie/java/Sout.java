package com.xie.java;

/**
 * @author xie4ever
 * @date 2021/5/12 18:02
 */
public class Sout {

    static Sout st = new Sout();

    static {
        System.out.println("1");
    }

    {
        System.out.println("2");
    }

    Sout() {
        System.out.println("3");
        System.out.println("a=" + a + ",b=" + b);
    }

    public static void fun() {
        System.out.println("4");
    }

    int a = 110;
    static int b = 112;

    public static void main(String[] args) {
        fun();
    }
}
