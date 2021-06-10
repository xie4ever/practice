package com.xie.exam;

import java.util.Scanner;

/**
 * @author xie4ever
 * @date 2021/5/23 0:09
 */
public class Change {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        while (scanner.hasNextLine()) {
            String str = scanner.nextLine();
            int num = Integer.parseInt(str.substring(2), 16);
            System.out.println(num);
        }
    }
}
