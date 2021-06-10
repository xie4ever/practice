package com.xie.nowcoder.output;

import java.util.*;

/**
 * @author xie4ever
 * @date 2021/5/16 21:33
 */
public class MyExam2 {

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        while (scanner.hasNextLine()) {
            String first = scanner.nextLine();
            String second = scanner.nextLine();
            HashMap<Character, Integer> map = new HashMap<>();
            HashMap<Character, Integer> existed = new HashMap<>();
            char[] chars = first.toCharArray();
            for (char aChar : chars) {
                map.put(aChar, 0);
            }
            List<Character> list = new ArrayList<>();
            char[] secondChars = second.toCharArray();
            for (char secondChar : secondChars) {
                if (existed.containsKey(secondChar)) {
                    continue;
                }
                existed.put(secondChar, 0);
                if (map.containsKey(secondChar)) {
                    list.add(secondChar);
                }
            }
            Collections.sort(list);
            StringBuilder stringBuilder = new StringBuilder();
            for (char l : list) {
                stringBuilder.append(l);
            }
            System.out.println(stringBuilder.toString());
        }
    }
}
