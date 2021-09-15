from readability import Document
from lxml import html
import requests

def getPlain(url):
    """Get plain text content of URL

    :param url: URL to fetch content.
    """
    resp = requests.get(url)
    doc = Document(resp.text)
    el = html.fromstring(doc.summary())
    return el.text_content().translate(str.maketrans('', '', '\n\r\t'))

if __name__ == "__main__":
    import sys
    for url in sys.argv[1:]:
        try:
            print("OK\t%s\t%s" % (url, getPlain(url)))
        except Exception as e:
            print("NG\t%s\t%s" % (url, e))
