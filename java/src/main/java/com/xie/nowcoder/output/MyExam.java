package com.xie.nowcoder.output;

import java.util.HashMap;
import java.util.Scanner;

/**
 * @author xie4ever
 * @date 2021/5/16 20:58
 */
public class MyExam {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String str = scanner.nextLine();
        String[] strs = str.split(",");
        long max = 0;
        HashMap<Character, Integer> hashMap;
        for (int i = 0; i < strs.length; i++) {
            hashMap = new HashMap<>();
            String base = strs[i].toLowerCase();
            char[] chars = base.toCharArray();
            for (char aChar : chars) {
                hashMap.put(aChar, 0);
            }
            for (int j = i + 1; j < strs.length; j++) {
                String curr = strs[j].toLowerCase();
                char[] currChars = curr.toCharArray();
                boolean flag = true;
                for (char currChar : currChars) {
                    if (hashMap.containsKey(currChar)) {
                        flag = false;
                        break;
                    }
                }
                if (!flag) {
                    continue;
                }

                long num = base.length() * curr.length();
                if (num > max) {
                    max = num;
                }
            }
        }
        System.out.println(max);
    }
}
