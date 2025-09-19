gocat:
	go build ./...

all: single multiple stdin n b nb

single: gocat
	./gocat test.txt

multiple: gocat
	./gocat  test.txt test2.txt

stdin: gocat
	head -n1 test.txt | ./gocat -

n: gocat
	./gocat -n test.txt

b: gocat
	./gocat -b test.txt

nb: gocat
	./gocat -b -n test.txt

clean:
	rm gocat
