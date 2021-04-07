package com.xie.aqs;

import java.util.concurrent.locks.ReadWriteLock;
import java.util.concurrent.locks.ReentrantReadWriteLock;

/**
 * @author xie4ever
 * @date 2021/3/31 9:20
 */
public class TestAQS {
    ReadWriteLock lock = new ReentrantReadWriteLock();

    public void testRead() {
        new Thread(() -> {
            while (true) {
                try {
                    lock.readLock().lock();
                    System.out.println("read a lot");
                    try {
                        Thread.sleep(1000);
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                } finally {
                    lock.readLock().unlock();
                }
            }
        }).start();
    }

    public void testWrite() {
        new Thread(()->{
            try {
                lock.writeLock().lock();
                System.out.println("write a lot");
                try {
                    Thread.sleep(3000);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            } finally {
                lock.writeLock().unlock();
            }
        }).start();
    }

    public static void main(String[] args) {
        TestAQS testAQS = new TestAQS();
        testAQS.testWrite();
        testAQS.testRead();
    }
}
