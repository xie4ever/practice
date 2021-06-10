package com.xie.nowcoder.output;

import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

/**
 * @author xie4ever
 * @date 2021/5/16 22:04
 */
public class MyExam3 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String firstLine = scanner.nextLine();
        scanner.hasNextLine();
        String[] strings = firstLine.split(" ");
        int n = strings.length;
        int[][] arrs = new int[n][n];
        for (int i = 0; i < n; i++) {
            arrs[0][i] = Integer.valueOf(strings[i]);
        }
        for (int i = 1; i < n; i++) {
            for (int j = 0; j < n; j++) {
                arrs[i][j] = scanner.nextInt();
            }
        }

        int num = n;
        Map<String, Integer> map = new HashMap<>();
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (i == j) {
                    continue;
                }
                if (arrs[i][j] == 1) {
                    String fuckyounow = fuckyou(i, j);
                    String fuckmenow = fuckme(i, j);
                    if (map.containsKey(fuckyounow)) {
                        continue;
                    }
                    map.put(fuckmenow, 0);
                    num--;
                }
            }
        }

        System.out.println(num);
    }

    public static String fuckyou(int i, int j) {
        return j + "-" + i;
    }

    public static String fuckme(int i, int j) {
        return i + "-" + j;
    }
}
