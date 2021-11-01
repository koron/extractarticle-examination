default: test4_out.tsv test4_summary.tsv

test4: test4.go
	go build ./test4.go

test4_out.tsv: test4 ../dataset/*.html
	./test4 ../dataset/*.html > $@

test4_summary.tsv: test4_out.tsv
	cut -f 1,2,3,4,5,6,7 < $< > $@

clean:
	rm -f test4

distclean:
	rm -f test4_out.tsv
	rm -f test4_summary.tsv
