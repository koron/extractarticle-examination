package net.kaoriya.extractarticle;

import java.util.Map;
import java.util.HashMap;

public class Bigram {
    public final Map<String, Integer> index = new HashMap<String, Integer>();

    public Bigram(String str) {
        for (int i = 0; i < str.length() - 1; i++) {
            var sub = str.substring(i, i+2);
            int cnt = 0;
            if (index.containsKey(sub)) {
                cnt = index.get(sub);
            }
            index.put(sub, cnt + 1);
        }
    }

    public int count() {
        return index.size();
    }
}
