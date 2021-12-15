package net.kaoriya.extractarticle;

import java.io.File;
import java.io.FileInputStream;
import java.io.InputStreamReader;

import net.dankito.readability4j.Readability4J;
import org.jsoup.Jsoup;

public class App4 {
    public static void main(String[] args) throws Exception {
        var dir = new File("../dataset");
        for (var f : dir.listFiles()) {
            if (f.isDirectory() || !f.getName().endsWith(".html")) {
                continue;
            }
            var r = ArticleExtractor.extract(f);
            System.out.println(String.format("OK\t%s\t%g\t%d\t%d", f.getName(), r.score, r.text.length(), r.desc.length()));
        }
    }
}
