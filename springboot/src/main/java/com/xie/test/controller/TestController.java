package com.xie.test.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.time.Duration;

/**
 * @author xie4ever
 * @date 2021/3/23 23:23
 */
@RestController
@RequestMapping("test")
public class TestController {

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
}
