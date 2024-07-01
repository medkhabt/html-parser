# HTML Parser

I am basing this implementation with golang on this [doc](https://www.w3.org/TR/2011/WD-html5-20110113/parsing.html). 

Currently I am in the tokenization phase. 

the lexer has many states. I am building a map for clear global idea on the stats and also to track where i am at in the implementation, so it's certain that there is no state that is not present in the map and that is implemented in this repo.

here is the current map 
![map_01_july_2024](.resources/imgs/htmlparser_01_Juli_2024.png)
[UPDATED 01/07/2024]
