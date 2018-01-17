#!/bin/bash

#build all the asciidoc
ls | grep ".adoc" | xargs -n 1 asciidoctor -a linkcss

#remove the css file that comes from the asciidoctor command
rm asciidoctor.css

#Removes the entire <head> section
sed -i '/<head>/,/<\/head>/d' *.html

#define document at the top
sed -i '1i{{define "document"}}' *.html

#define end at the bottom
sed -i -e "\$a{{end}}" *.html




