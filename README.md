# gRPC Animales

Este proyecto implementa un servidor gRPC para gestionar información sobre animales. Permite realizar consultas, añadir nuevos registros, y filtrar animales en base a atributos como el nombre, tipo, o número de patas. El sistema está construido en Go y utiliza una base de datos SQL Server para almacenar los datos.

## Descripción General

El servicio gRPC maneja las siguientes operaciones relacionadas con los animales:

- **Obtener información de un animal específico** por su nombre.
- **Listar todos los animales** disponibles en la base de datos.
- **Añadir nuevos animales** al sistema mediante una transmisión de datos en streaming.
- **Filtrar animales por tipo**, devolviendo una lista de animales que coinciden con el tipo especificado.

## Requisitos Previos

- **Lenguaje**: Go
- **Base de datos**: SQL Server
- **Protocolo**: gRPC
- **Dependencias**:
  - `github.com/denisenkom/go-mssqldb`
  - `github.com/joho/godotenv`
  - `google.golang.org/grpc`

## Estructura de Protocolo gRPC

El archivo `.proto` define los servicios y mensajes utilizados:

```proto
syntax = "proto3";

package animales;

option go_package = "Animales-go/proto;proto";

service animalesservice {
    rpc GetAnimalesInfo (AnimalRequest) returns (AnimalResponse);
    rpc GetAnimalesList( Empty ) returns (stream AnimalResponse);
    rpc AddAnimales(stream NewAnimalRequest) returns (AddAnimalResponse);
    rpc GetAnimalesByType(stream AnimalTypeRequest) returns (stream AnimalResponse);
}

message AnimalRequest {
    string name = 1;
}

message AnimalResponse {
    string Name = 1;
    string Type = 2;
    int32 NumPatas = 3;
}

message NewAnimalRequest {
    string Name = 1;
    string Type = 2;
    int32 NumPatas = 3;
}

message AddAnimalResponse {
    int32 Count = 1;
}

message AnimalTypeRequest {
    string Type = 1;
}

message Empty {}
