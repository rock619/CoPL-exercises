generate: clean
	antlr -Dlanguage=Go -visitor -package parser -o parser *.g4

clean:
	rm -rf parser/*
