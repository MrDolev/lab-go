# Jaro–Winkler

This code implementation about [Jaro-Winkler](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance) is inspired by [Apache Commons Text][github-apache]. 

This code is not intended to be an official library, but only to compare existing implementations on text similarity to analyse performance and accuracy between different libraries.

The library used to check beanchmark and accuracy are;
- [matchr][github-matchr]
- [smetrics][github-smetrics]

# Project strcuture
Inside examples you can compile and execute the main program that is possible to investigate the benchmark, accuracy about three implementations about similarity.

# Run the benchmark and accuracy
On examples folder is present the main package that is possible to run the benchmark and accuracy and eventually adding more test case examples inspired from [Apache Commons Text][github-apache] and [github-matchr].

# Perform test
Inside in the internal package is possible to run the tests about similarity and distance.

[github-apache]: https://github.com/apache/commons-text/tree/master/src/main/java/org/apache/commons/text/similarity
[github-matchr]: https://github.com/antzucaro/matchr
[github-smetrics]: https://github.com/xrash/smetrics