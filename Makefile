run: check-modd-exists
	@modd -f ./.modd/modd.conf

check-modd-exists:
	@modd --version > /dev/null