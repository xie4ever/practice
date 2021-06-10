package com.xie.nowcoder.output;

import java.util.Scanner;

/**
 * @author xie4ever
 * @date 2021/5/15 1:00
 */
public class Test1 {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int x = scanner.nextInt();
        int y = scanner.nextInt();
        int[][] ints = new int[x][y];
        scanner.nextLine();
        for (int i = 0; i < x; i++) {
            String str = scanner.nextLine();
            String[] strs = str.split(" ");
            for (int j = 0; j < strs.length; j++) {
                ints[i][j] = Integer.valueOf(strs[j]);
            }
        }

        for (int[] anInt : ints) {
            for (int i : anInt) {
                System.out.print(i + " ");
            }
            System.out.println();
        }
    }
}
