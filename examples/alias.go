package main

type Name1 map[string]string
type Name2 map[string]string
type Alias = map[string]string

var simple map[string]string
var n1 Name1
var n2 Name2
var al Alias


simple = n1 // Possible
n1 = n2 // NOT POSSIBLE

al = n1 // Possible
