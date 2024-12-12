package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	pb "Animales-go/proto"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var db *sql.DB

type server struct {
	pb.UnimplementedAnimalesserviceServer
}

func (s *server) GetAnimalesInfo(ctx context.Context, req *pb.AnimalRequest) (*pb.AnimalResponse, error) {
	var name, ptype string
	var numPatas int

	query := "select * from animalsprd.Animal where Name like @Name"

	row := db.QueryRowContext(ctx, query, sql.Named("Name", req.Name))
	err := row.Scan(&name, &ptype, &numPatas)

	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.AnimalResponse{
				Name:     "Not Found",
				Type:     "Not Found",
				NumPatas: 0,
			}, nil

		}
		return nil, err
	}

	return &pb.AnimalResponse{
		Name:     name,
		Type:     ptype,
		NumPatas: int32(numPatas),
	}, nil
}

func (s *server) GetAnimalesList(req *pb.Empty, stream pb.Animalesservice_GetAnimalesListServer) error {
	query := "select * from animalsprd.Animal"
	rows, err := db.Query(query)

	if err != nil {
		log.Panic(err)
		return err
	}
	defer rows.Close()
	///////
	for rows.Next() {
		var name, ptype string
		var numPatas int

		if err := rows.Scan(&name, &ptype, &numPatas); err != nil {
			log.Panic(err)
			return err
		}

		if err := stream.Send(&pb.AnimalResponse{
			Name:     name,
			Type:     ptype,
			NumPatas: int32(numPatas),
		}); err != nil {
			log.Panic(err)
			return err

		}
	}

	return nil
}

func (s *server) AddAnimales(stream pb.Animalesservice_AddAnimalesServer) error {
	var count int32

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AddAnimalResponse{
				Count: count,
			})
		}

		if err != nil {
			log.Panic(err)
			return err
		}

		query := "insert into animalsprd.Animal (Name, Type, NumPatas) values (@Name, @Type, @NumPatas)"
		_, err = db.Exec(query,
			sql.Named("Name", req.Name),
			sql.Named("Type", req.Type),
			sql.Named("NumPatas", req.NumPatas))

		if err != nil {
			log.Panic(err)
			return err
		}

		count++
		log.Printf("Added %s", req.Name)
	}
}

func (s *server) GetAnimalesByType(stream pb.Animalesservice_GetAnimalesByTypeServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of stream")
			return nil
		}

		if err != nil {
			log.Panic(err)
			return err
		}

		query := "select * from animalsprd.Animal where lower(Type) = lower(@Type)"
		rows, err := db.Query(query, sql.Named("Type", req.Type))
		if err != nil {
			log.Panic(err)
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var name, ptype string
			var numPatas int

			if err := rows.Scan(&name, &ptype, &numPatas); err != nil {
				log.Panic(err)
				return err
			}

			if err := stream.Send(&pb.AnimalResponse{
				Name:     name,
				Type:     ptype,
				NumPatas: int32(numPatas),
			}); err != nil {
				log.Panic(err)
				return err
			}
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro loading .env  file")
	}

	s := os.Getenv("DB_SERVER")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, s, port, database)

	log.Println("esta es la URL: ", connString)

	db, err = sql.Open("sqlserver", connString)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Connected to database")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAnimalesserviceServer(grpcServer, &server{})

	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
		log.Println("Starting health check on port 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	log.Println("Starting server on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}

}
