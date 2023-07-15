## Important commands:

### To generate mock:
`mockery --dir=postgres<src directory> --output=mocks<o/p directory> --outpkg=mock<o/p package> --all`

### To run tests
`go test -v ./...`

### Create and Run migrations using golang-migrate
`migrate create -ext sql -dir migrations -seq create_bill_table`

`migrate -path migrations -database "postgresql://user:password@localhost:5432/internet_bills?sslmode=disable" up`
