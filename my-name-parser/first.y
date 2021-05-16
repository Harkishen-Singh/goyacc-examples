%{

package main

import "fmt"

var firstname, lastname bool

%}

%union{
    s string
}

%type<s> sentence keyword either

%token<s> HARKISHEN SINGH RANDOM

%%

result: sentence                {
        if firstname && lastname {
            fmt.Println("sentence contains both, my first and last name")
        } else {
            fmt.Println("name not present")
        }
    }

sentence: keyword               { $$ = $1 }
    |   sentence keyword        { $$ = fmt.Sprintf("%s %s", $1, $2)  }
    |   keyword sentence        { $$ = fmt.Sprintf("%s %s", $1, $2)  }
    ;

keyword:    HARKISHEN SINGH     { $$ = fmt.Sprintf("%s %s", $1, $2) }
    |   RANDOM                  
    |   either RANDOM           { $$ = $1 }
    |   RANDOM either           { $$ = $2 }
    |   either                  { $$ = $1 }
    ;

either:     HARKISHEN           { $$ = $1; firstname = true; }
    |   SINGH                   { $$ = $1; lastname = true; }
    ;

%%
