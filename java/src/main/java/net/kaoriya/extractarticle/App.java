package net.kaoriya.extractarticle;

import java.net.URL;

import de.l3s.boilerpipe.BoilerpipeProcessingException;
import de.l3s.boilerpipe.extractors.ArticleExtractor;
import de.l3s.boilerpipe.extractors.ArticleSentencesExtractor;
import de.l3s.boilerpipe.extractors.DefaultExtractor;
import de.l3s.boilerpipe.extractors.LargestContentExtractor;

public class App {
    final static String[] urlList = new String[] {
        "https://news.yahoo.co.jp/articles/acba39a56e3a15a628b2c6e7c288f103222ce86c",
        "https://news.yahoo.co.jp/pickup/6404492",
        "https://www.asahi.com/articles/ASP9G5DG1P9GUTIL011.html",
        "https://www.asahi.com/articles/ASP9G66HWP9GUCLV00D.html",
        "https://www.asahi.com/ads/springvalley202109/",
        "https://www.yomiuri.co.jp/economy/20210915-OYT1T50049/",
        "https://www.yomiuri.co.jp/culture/20210914-OYT1T50344/",
        "https://yab.yomiuri.co.jp/adv/chuo/opinion/20210311.php",
        "https://www.sankei.com/article/20210915-HGJUDSPJKNBCXMXGCO4I5VTBGY/",
        "https://www.sankei.com/article/20210911-E523ZF5SKJI7BBYWAZG6WYIQXM/",
        "https://mainichi.jp/articles/20210915/k00/00m/040/074000c",
        "https://mainichi.jp/articles/20210914/dde/012/040/017000c",
        "https://mainichi.jp/sp/sekai/",
        "https://japanese.engadget.com/jp-pr-oppo-reno-5-a-020005713.html",
        "https://nazology.net/archives/95980",
    };

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

    public static void extract(String urlStr) throws Exception {
        var extractor = ArticleExtractor.getInstance();
        //var extractor = LargestContentExtractor.getInstance();
        URL url = new URL(urlStr);
        try {
            String text = extractor.getText(url);
            text = singlePlainText(text);
            System.out.print(String.format("OK\t%s\t%s\n", urlStr, text));
        } catch (BoilerpipeProcessingException e) {
            System.out.print(String.format("NG\t%s\t%s\n", urlStr, e.toString()));
        }
    }

    public static void main(String[] args) throws Exception {
        for (String s : urlList) {
            extract(s);
        }
    }
}
