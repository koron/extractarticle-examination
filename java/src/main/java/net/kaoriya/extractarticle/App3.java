package net.kaoriya.extractarticle;

import java.io.File;
import java.io.FileInputStream;
import java.io.InputStreamReader;

import net.dankito.readability4j.Readability4J;
import org.jsoup.Jsoup;

public class App3 {

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
        for (var f : dir.listFiles()) {
            if (f.isDirectory() || !f.getName().endsWith(".html")) {
                continue;
            }
            var doc = Jsoup.parse(f, "UTF-8");
            var r4j = new Readability4J("", doc);
            var article = r4j.parse();
            String text = singlePlainText(article.getTextContent());
            System.out.println(String.format("OK\t%s\t%s", f.getName(), text));
        }
    }
}
