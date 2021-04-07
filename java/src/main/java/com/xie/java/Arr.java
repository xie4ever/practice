package com.xie.java;

/**
 * @author xie4ever
 * @date 2021/3/23 13:44
 */
public class Arr {

    private static void test1() {
        int[] arr = new int[]{0, 0};
        int a = arr[1];
        arr[1] = 1;
        System.out.println(a); // 1 × 0 √
        System.out.println(arr[1]);
    }

    private static void test2() {
        int[][] arr2 = new int[2][2];

        int[] tmp0 = new int[]{1};
        int[] tmp1 = new int[]{2, 3, 4};

        arr2[0] = tmp0;
        arr2[1] = tmp1;

        int[][] arr3 = new int[3][];
        arr3[0] = arr2[0];
        arr3[1] = arr2[1];

        System.out.println(arr3[0] == tmp0);
        System.out.println(arr3[1] == tmp1);

        arr2[0] = new int[]{1, 2};
        arr2[1] = new int[]{2, 3, 4};

        System.out.println(arr3[0] == tmp0);
        System.out.println(arr3[1] == tmp1);

        System.out.println(arr3[0] == arr2[0]);
        System.out.println(arr3[1] == arr2[1]);
    }

    public static void main(String[] args) {
        test2();
    }
}
