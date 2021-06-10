package com.xie.sword2offer;

/**
 * @author xie4ever
 * @date 2021/5/25 0:24
 */
public class JZ2 {

    public String replaceSpace (String s) {
        StringBuilder stringBuilder = new StringBuilder();
        char[] chars = s.toCharArray();
        for (int i = 0; i < chars.length; i++) {
            if (chars[i] == ' ') {
                stringBuilder.append("%20");
            } else {
                stringBuilder.append(chars[i]);
            }
        }
        return stringBuilder.toString();
    }
}
