package com.xie.deadlock;

import org.junit.jupiter.api.Test;

import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

/**
 * @author xie4ever
 * @date 2021/5/26 23:38
 */
public class TestDeadlock {

    @Test
    public void test() throws InterruptedException {
        Lock lock1 = new ReentrantLock();
        Lock lock2 = new ReentrantLock();

        class t1 extends Thread {
            @Override
            public void run() {
                for (int i = 0; i < 1000; i++) {
                    lock1.lock();
                    System.out.println("t1 lock1");
                    lock2.lock();
                    System.out.println("t1 lock2");
                    lock2.unlock();
                    lock1.unlock();
                }
            }
        }

        class t2 extends Thread {
            @Override
            public void run() {
                for (int i = 0; i < 1000; i++) {
                    lock2.lock();
                    System.out.println("t2 lock2");
                    lock1.lock();
                    System.out.println("t2 lock1");
                    lock1.unlock();
                    lock2.unlock();
                }
            }
        }

        t1 thread1 = new t1();
        t2 thread2 = new t2();
        thread1.start();
        thread2.start();
        thread1.join();
        thread2.join();
    }

    @Test
    public void test2() throws InterruptedException {
        Object ob1 = new Object();
        Object ob2 = new Object();

        class t1 extends Thread {
            @Override
            public void run() {
                while (true) {
                    synchronized (ob1) {
                        synchronized (ob2) {
                            System.out.println("t1");
                        }
                    }
                }
            }
        }

        class t2 extends Thread {
            @Override
            public void run() {
                while (true) {
                    synchronized (ob2) {
                        synchronized (ob1) {
                            System.out.println("t2");
                        }
                    }
                }
            }
        }

        t1 thread1 = new t1();
        t2 thread2 = new t2();
        thread1.start();
        thread2.start();
        thread1.join();
        thread2.join();
    }

    @Test
    public void test3() throws InterruptedException {

        class t1 extends Thread {
            @Override
            public void run() {
                int num = 0;
                while (!this.isInterrupted()) {
                    num++;
                    try {
                        Thread.sleep(100);
                    } catch (InterruptedException e) {
                       return;
                    }
                }
                System.out.println(num);
            }
        }

        t1 thread1 = new t1();
        thread1.start();
        Thread.sleep(1000);
        thread1.interrupt();
    }
}
