prep:
	echo $(DAY) > day.txt
	cp -r template $(DAY)

input: DAY=$(shell cat day.txt)
input: COOKIE=$(shell cat cookie.txt)
input: URL_DAY=$(shell echo $(DAY) | sed 's/^0*//')
input:
	curl --location --request GET "https://adventofcode.com/2023/day/$(URL_DAY)/input" \
		--header "Cookie: $(COOKIE)" --output "$(DAY)/testdata/input.txt"

test: DAY = $(shell cat day.txt)
test:
	go test -v $(DAY)/*.go