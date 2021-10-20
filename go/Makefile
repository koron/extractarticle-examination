default: test4_out.txt test4_summary.txt test4_excerpt.txt test4_body.txt

test4: test4.go
	go build ./test4.go

test4_out.txt: test4 ../dataset/*.html
	./test4 ../dataset/*.html > $@

test4_summary.txt: test4_out.txt
	cut -f 1,2,3,4,5 < $< > $@

test4_excerpt.txt: test4_out.txt
	cut -f 2,6 < $< > $@

test4_body.txt: test4_out.txt
	cut -f 2,7 < $< > $@

clean:
	rm -f test4

distclean:
	rm -f test4_out.txt
	rm -f test4_summary.txt
	rm -f test4_title.txt
