.PHONY: tag patch minor major

# Deprecated tag functiontag:
tag:
	@echo "The 'tag' function is deprecated. Please use 'make patch', 'make minor', or 'make major' instead."

# Run patch version increment
patch:
	@bash executepatch.sh

# Run minor version increment
minor:
	@bash executeminor.sh

# Run major version increment
major:
	@bash executemajor.sh