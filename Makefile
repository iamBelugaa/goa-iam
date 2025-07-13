gen-goa:
	@goa gen github.com/iamBelugaa/goa-iam/internal/design
	@goa gen github.com/iamBelugaa/goa-iam/internal/services/authsvc/design -o ./gen/auth
	@goa gen github.com/iamBelugaa/goa-iam/internal/services/usersvc/design -o ./gen/user