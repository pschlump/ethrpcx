

all:

b0:
	git checkout -b fix-readme-example

# make changes to readme.md file
# git add .
b1:
	git commit -m 'Fixed the example at the bottom of readme to compile and run'
	git push origin readme-example
	echo now go to github and create a PR
