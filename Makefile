.PHONY: test-anagram
test_anagram:
	cd anagram && go test ./...

.PHONY: test-brokennode
test-brokennode:
	cd brokennode && go test ./...

.PHONY: test-buildworld
test-buildworld:
	cd buildworld && go test ./...

.PHONY: test-chess
test-chess:
	cd chess && go test ./...

.PHONY: test-floyd
test-floyd:
	cd floyd && go test ./...

.PHONY: test-functionfrequency
test-functionfrequency:
	cd functionfrequency && go test ./...

.PHONY: test-jaro
test-jaro:
	cd jaro && go test ./...

.PHONY: test-lastlettergame
test-lastlettergame:
	cd lastlettergame && go test ./...

.PHONY: test-mergesort
test-mergesort:
	cd mergesort && go test ./...

.PHONY: test-missingnumbers
test-missingnumbers:
	cd missingnumbers && go test ./...

.PHONY: test-node_degree
test-node_degree:
	cd node_degree && go test ./...

.PHONY: test-reverseparentheses
test-reverseparentheses:
	cd reverseparentheses && go test ./...

.PHONY: test-romannumerals
test-romannumerals:
	cd romannumeraks && go test ./...

.PHONY: test-secretmessage
test-secretmessage:
	cd secretmessage && go test ./...

.PHONY: test-shorthash
test-shorthash:
	cd shorthash && go test ./...

.PHONY: test-sumdecimal
test-sumdecimal:
	cd sumdecimal && go test ./...

.PHONY: test-warriors
test-warriors:
	cd warriors && go test ./...

.PHONY: test-wordladder
test-wordladder:
	cd wordladder && go test ./...
