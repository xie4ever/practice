package com.xie.test.controller;

import org.redisson.api.RedissonClient;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.time.Duration;
import java.util.Objects;
import java.util.UUID;

/**
 * @author xie4ever
 * @date 2021/3/23 23:23
 */
@RestController
@RequestMapping("test")
public class TestController {

    @Autowired
    private RedissonClient redissonClient;

    @Autowired
    private StringRedisTemplate stringRedisTemplate;

    @RequestMapping("/test")
    public String test() {
        return "test";
    }

    @PutMapping
    @RequestMapping("/put")
    public String put(@RequestParam(value = "value", required = true) String value) {
        String key = "test";
        String test = stringRedisTemplate.opsForValue().get(key);
        if (test == null || test.isEmpty()) {
            stringRedisTemplate.opsForValue().set(key, value, Duration.ofSeconds(10));
            return "cache";
        }
        return test;
    }

    public void test1() {
        synchronized (this) {
            int stock = Integer.parseInt(Objects.requireNonNull(stringRedisTemplate.opsForValue().get("stock")));
            if (stock > 0) {
                int realStock = stock - 1;
                stringRedisTemplate.opsForValue().set("stock", realStock + "");
                System.out.println("success");
            } else {
                System.out.println("fail");
            }
        }
    }

    public void test2() {
        boolean lock = stringRedisTemplate.opsForValue().setIfAbsent("lockKey", "lockValue");
        if (!lock) {
            System.out.println("lock fail");
            return;
        }

        int stock = Integer.parseInt(Objects.requireNonNull(stringRedisTemplate.opsForValue().get("stock")));
        if (stock > 0) {
            int realStock = stock - 1;
            stringRedisTemplate.opsForValue().set("stock", realStock + "");
            System.out.println("success");
        } else {
            System.out.println("fail");
        }

        stringRedisTemplate.delete("lockKey");
    }

    public void test3() {
        boolean lock = stringRedisTemplate.opsForValue().setIfAbsent("lockKey", "lockValue");
        if (!lock) {
            System.out.println("lock fail");
            return;
        }

        try {
            int stock = Integer.parseInt(Objects.requireNonNull(stringRedisTemplate.opsForValue().get("stock")));
            if (stock > 0) {
                int realStock = stock - 1;
                stringRedisTemplate.opsForValue().set("stock", realStock + "");
                System.out.println("success");
            } else {
                System.out.println("fail");
            }
        } finally {
            stringRedisTemplate.delete("lockKey");
        }
    }

    public void test4() {
        boolean lock = stringRedisTemplate.opsForValue().setIfAbsent("lockKey", "lockValue");
        stringRedisTemplate.expire("lockKey", Duration.ofSeconds(30));
        if (!lock) {
            System.out.println("lock fail");
            return;
        }

        try {
            int stock = Integer.parseInt(Objects.requireNonNull(stringRedisTemplate.opsForValue().get("stock")));
            if (stock > 0) {
                int realStock = stock - 1;
                stringRedisTemplate.opsForValue().set("stock", realStock + "");
                System.out.println("success");
            } else {
                System.out.println("fail");
            }
        } finally {
            stringRedisTemplate.delete("lockKey");
        }
    }

    public void test5() {
        boolean lock = stringRedisTemplate.opsForValue().setIfAbsent("lockKey", "lockValue", Duration.ofSeconds(30));
        if (!lock) {
            System.out.println("lock fail");
            return;
        }

        try {
            int stock = Integer.parseInt(Objects.requireNonNull(stringRedisTemplate.opsForValue().get("stock")));
            if (stock > 0) {
                int realStock = stock - 1;
                stringRedisTemplate.opsForValue().set("stock", realStock + "");
                System.out.println("success");
            } else {
                System.out.println("fail");
            }
        } finally {
            stringRedisTemplate.delete("lockKey");
        }
    }

    public void test6() {
        String clientID = UUID.randomUUID().toString();
        boolean lock = stringRedisTemplate.opsForValue().setIfAbsent("lockKey", clientID, Duration.ofSeconds(30));
        if (!lock) {
            System.out.println("lock fail");
            return;
        }

        try {
            int stock = Integer.parseInt(Objects.requireNonNull(stringRedisTemplate.opsForValue().get("stock")));
            if (stock > 0) {
                int realStock = stock - 1;
                stringRedisTemplate.opsForValue().set("stock", realStock + "");
                System.out.println("success");
            } else {
                System.out.println("fail");
            }
        } finally {
            if (clientID.equals(stringRedisTemplate.opsForValue().get("lockKey"))) {
                stringRedisTemplate.delete("lockKey");
            }
        }
    }

    public void test7() {
        String clientID = UUID.randomUUID().toString();
        boolean lock = stringRedisTemplate.opsForValue().setIfAbsent("lockKey", clientID, Duration.ofSeconds(30));
        if (!lock) {
            System.out.println("lock fail");
            return;
        }

        try {
            int stock = Integer.parseInt(Objects.requireNonNull(stringRedisTemplate.opsForValue().get("stock")));
            if (stock > 0) {
                int realStock = stock - 1;
                stringRedisTemplate.opsForValue().set("stock", realStock + "");
                System.out.println("success");
            } else {
                System.out.println("fail");
            }
        } finally {
            redissonClient.
        }
    }
}
