%{

package main

import "fmt"

// This package parses phone numbers according to Indian standards.

%}

%union{
    digit string
    sym   byte
    random string
    str    string
}

%type<str>    digit_set digit_pair number phone_number

%token<digit>   DIGIT
%token<sym>     SYMBOL
%token<random>  RANDOM_STRING

%%

result:                 phone_number        {
        fmt.Println("phone number is valid according to Indian standards")
    }
    ;

phone_number:           SYMBOL digit_pair number       { $$ = concat(string($1), $2) }
    |                   number              { $$ = $1 }
    ;

number:                 digit_set digit_set { $$ = concat($1, $2) }
    ;

digit_pair:             DIGIT DIGIT         { $$ = concat($1, $2) }
    |                   DIGIT               { $$ = $1 }
    ;

digit_set:              DIGIT DIGIT DIGIT DIGIT DIGIT       { $$ = concat($1, $2, $3, $4, $5) }
    ;

%%
