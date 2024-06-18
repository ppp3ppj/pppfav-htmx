migrate-create:
	echo "Creating migration wywy_db"
	migrate create -ext sql -dir db/migrations -seq pppfav_db

migrate-up:
	echo "Migrating-up database"
	migrate -source file://db/migrations -database 'postgres://pppwywy:123456@localhost:4444/pppfav_db?sslmode=disable' -verbose up

migrate-down:
	echo "Migrating-down database"
	migrate -source file://db/migrations -database 'postgres://pppwywy:123456@localhost:4444/pppfav_db?sslmode=disable' -verbose down
