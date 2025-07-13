# Remove all generated files.
clean:
	@echo "Cleaning generated files..."
	@rm -rf ./gen
	@echo "Done."

# Generate Goa design packages.
gen-goa: clean
	@echo "Generating Goa code..."
	@goa gen github.com/iamBelugaa/goa-iam/internal/design
	@goa gen github.com/iamBelugaa/goa-iam/internal/services/authsvc/design -o ./gen/auth
	@goa gen github.com/iamBelugaa/goa-iam/internal/services/usersvc/design -o ./gen/user
	@echo "Goa code generation complete."

.PHONY: clean gen-goa