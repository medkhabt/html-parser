# HTML Parser

I am basing this implementation with golang on this [doc](https://www.w3.org/TR/2011/WD-html5-20110113/parsing.html). 

Currently I am in the tokenization phase. 

the lexer has many states. I am building a map for clear global idea on the stats and also to track where i am at in the implementation, so it's certain that there is no state that is not present in the map and that is implemented in this repo.

here is the current map 
![map_02_july_2024](.resources/imgs/htmlparser_02_Juli_2024.png)
I will uplaod new parts as new images, it become impossible to import the entire map, so i had to slice just the new parts :) 

update 03/07/2024 (expend on the attribute name and value)

![map_slice_03_july_2024](.resources/imgs/htmlparser_03_Juli_2024.png)

update 04/07/2024 (fix some mistakes on the map and add better placement to the after attribute value quoted state) 
![map_slice_04_july_2024](.resources/imgs/htmlparser_04_Juli_2024.png)

[UPDATED 04/07/2024] (CURRENTLY ON HOLD, Wanted to make a vimscript to get TODO task in a nice way ^^')
