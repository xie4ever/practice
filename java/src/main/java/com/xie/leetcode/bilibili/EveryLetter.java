package com.xie.leetcode.bilibili;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

/**
 * @author xie4ever
 * @date 2021/3/15 16:42
 */
public class EveryLetter {

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        String str;
        while ((str = br.readLine()) != null) {
            String num = str.split(" ")[0];
            int k = Integer.parseInt(num);
            String s = str.substring(num.length() + 1);
            System.out.println(appearFirstK(k, s));
        }
    }

    private static String appearFirstK(int k, String s) {
        char[] chars = s.toCharArray();
        int[] hash = new int[200];
        for (char c : chars) {
            hash[c]++;
        }
        for (char c : chars) {
            if (hash[c] == 1 && k == 1) {
                return "[" + c + "]";
            } else if (hash[c] == 1) {
                k--;
            }
        }
        return "Myon~";
    }
}
