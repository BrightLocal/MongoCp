package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/BrightLocal/MongoCp/copier"
	"github.com/BrightLocal/MongoCp/dsn"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}
	fromDSN, toDSN := dsn.Parse(os.Args[1]), dsn.Parse(os.Args[1])
	cp := copier.New(connect(fromDSN), connect(toDSN))
	if err := cp.Copy(fromDSN, toDSN); err != nil {
		log.Fatalf("Error copying: %s", err)
	}
}

func usage() {
	const usage = `%s <from_dsn> <to_dsn>
<dsn> is "[username[:password]@]host[:port]/database[?flags]"
<flags> are key=value pairs separated by "&"
Supported flags:
auth -- authentication method, default "MONGODB-CR" 
`
	fmt.Printf(usage, os.Args[0])
}

func connect(dsn dsn.DSN) *mongo.Client {
	opts := &options.ClientOptions{}
	authMechanism := "MONGODB-CR"
	if am := dsn.GetExtra("auth"); am != "" {
		authMechanism = am
	}
	if dsn.UserName != "" {
		opts = opts.SetAuth(options.Credential{
			AuthMechanism: authMechanism,
			AuthSource:    dsn.Database,
			Username:      dsn.UserName,
			Password:      dsn.Password,
			PasswordSet:   dsn.Password != "",
		})
	}
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatalf("Error connecting to %q: %s", dsn.HostName, err)
	}
	return client
}
