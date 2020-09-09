# Simple Full-Text Search for The Building Coder

I discovered a nice article by Artem Krylysov
suggesting [let's build a full-text search engine](https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine)
using the [Go programming language](https://golang.org),
with an accompanying [simplefts GitHub repository](https://github.com/akrylysov/simplefts) sharing the final result.

That prompted me to dabble a bit with Go and use it to implement full text search functionality for The Building coder blog posts as well.

The work in progress is available from my [tbcfts GitHub repository](https://github.com/jeremytammik/tbcfts).

[The Building Coder blog source `tbc`](https://github.com/jeremytammik/tbc) is available from GitHub as well, so you have all you need to play with it yourself, if you please.

Here is a sample run searching for the word 'dabble', which includes today's draft post, still lacking the public URL:

<pre>
/a/src/go/tbcfts $ ./tbcfts -q "dabble"
2020/09/09 10:31:40 Starting tbcfts, p=/a/doc/revit/tbc/git/a, q=dabble
2020/09/09 10:31:41 Loaded 1863 documents in 377.397917ms
2020/09/09 10:31:44 Indexed 1863 documents in 2.876775333s
2020/09/09 10:31:44 Search for 'dabble' found 5 documents in 9.703µs
2020/09/09 10:31:44 582 [Wiki API Help, View Event and Structural Material Type](http://thebuildingcoder.typepad.com/blog/2011/05/wiki-api-help-view-event-and-structural-material-type.html)
2020/09/09 10:31:44 906 [Export Wall Parts Individually to DXF](http://thebuildingcoder.typepad.com/blog/2013/03/export-wall-parts-individually-to-dxf.html)
2020/09/09 10:31:44 961 [Super Insane MP3 and Songbird Playlist Exporter](http://thebuildingcoder.typepad.com/blog/2013/06/super-insane-mp3-and-songbird-playlist-exporter.html)
2020/09/09 10:31:44 1008 [Open MEP Connector Warning](http://thebuildingcoder.typepad.com/blog/2013/08/open-mep-connector-warning.html)
2020/09/09 10:31:44 1863 [Optimising Parameters and Full-Text Search](http thebuildingcoder.typepad.com not yet published)
</pre>

According to `wc`, the current blog post HTML source consists of 355233 lines, 2230690 words and 20676311 characters, including markup.

As you can see, loading the documents and storing their body text in memory costs ca. 400 ms.

The indexing is costly, clocking in at ca. 3 seconds.

Once indexing is complete, the lookup is very fast, consuming just 10 microseconds.

Obviously, the next feature to address would be caching the index.

Another important enhancement would be to split the documents into smaller sections.

For instance, I could create much smaller and more targeted documents to index by using the `h4` tags that delimit individual sections of text within each blog post instead of retaining each blog post in its entirety as a single document.

Published 2020-09-09 by The Building Coder in the article
on [Optimising Parameters and Full-Text Search](https://thebuildingcoder.typepad.com/blog/2020/09/optimising-parameters-and-full-text-search.html).

## Todo

- Split documents into smaller sections based on `h4` tags instead of `h3`.
- Cache the index between runs.

## Author

Jeremy Tammik,
[The Building Coder](http://thebuildingcoder.typepad.com),
[Autodesk](http://www.autodesk.com)
[Forge Platform Development](http://forge.autodesk.com) and
[ADN](http://www.autodesk.com/adn)
[Open](http://www.autodesk.com/adnopen)

## License

This sample is licensed under the terms of the [MIT License](http://opensource.org/licenses/MIT).
Please see the [LICENSE](LICENSE) file for full details.
