package net.kaoriya.extractarticle;

import java.io.File;
import java.io.FileInputStream;
import java.io.InputStreamReader;

import de.l3s.boilerpipe.BoilerpipeProcessingException;
import de.l3s.boilerpipe.extractors.ArticleExtractor;
import de.l3s.boilerpipe.extractors.ArticleSentencesExtractor;
import de.l3s.boilerpipe.extractors.DefaultExtractor;
import de.l3s.boilerpipe.extractors.LargestContentExtractor;

public class App2 {

    public static String singlePlainText(String s) {
        StringBuilder b = new StringBuilder();
        boolean prev = false;
        for (char ch : s.toCharArray()) {
            if (ch == '　') {
                ch = ' ';
            }
            switch (ch) {
                case '\t':
                case '\n':
                case '\r':
                    prev = false;
                    continue;

                case ' ':
                    if (prev) {
                        continue;
                    }
                    prev = true;
                    break;

                default:
                    prev = false;
                    break;
            }
            b.append(ch);
        }
        return b.toString();
    }

    public static void main(String[] args) throws Exception {
        var dir = new File("../dataset");
        var ex = ArticleExtractor.getInstance();
        for (var f : dir.listFiles()) {
            if (f.isDirectory()) {
                continue;
            }
            if (!f.getName().endsWith(".html")) {
                continue;
            }
            try (var r = new InputStreamReader(new FileInputStream(f), "UTF-8")) {
                String text = singlePlainText(ex.getText(r));
                System.out.println(String.format("OK\t%s\t%s", f.getName(), text));
            }
        }
    }
}
