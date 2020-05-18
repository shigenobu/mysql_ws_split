# ws_split - MySQL UDF for split.

### about

This is MySQL User Defined Function written by cgo.  
Split from url to json in mysql string arg.  

|arg|explain|
|---|-------|
|arg1|`pattern` regexp pattern string|
|arg2|`subject` base string|

### how to install

    $ ./build.sh

(notice)  

* require root privilege

### example

(simple)  

    MariaDB [(none)]> select json_detailed(ws_split(',', 'aa,bb,cc')) as array;
    +----------------------------------+
    | array                            |
    +----------------------------------+
    | [
        "aa",
        "bb",
        "cc"
    ] |
    +----------------------------------+

(pattern)  

    MariaDB [(none)]> select json_detailed(ws_split('[\\s,]+', 'hypertext language, programming')) as array;
    +--------------------------------------------------------+
    | array                                                  |
    +--------------------------------------------------------+
    | [
        "hypertext",
        "language",
        "programming"
    ] |
    +--------------------------------------------------------+
