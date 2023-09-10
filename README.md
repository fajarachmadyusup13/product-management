# Product Management

## Background
Project ini menyediakan REST API untuk pengelolaan produk pada, dimana user mampu menambahkan produk baru, serta mampu menampilkan daftar produk yang tersimpan dalam database, dengan tambahan fitur mengurutkan produk berdasarkan nama, harga, dan tanggal masuk produk.

## Architecture
Dalam Project ini di terapkan clean architecture, beberapa prinsip yang diterapkan diantaranya adalah:

- **The Dependency Rule**
    > This rule says that source code dependencies can only point inwards. Nothing in an inner circle can know anything at all about something in an outer circle. In particular, the name of something declared in an outer circle must not be mentioned by the code in the an inner circle. That includes, functions, classes. variables, or any other named software entity.
    
    Hal ini diterapkan dengan tujuan agar layer yang berada di paling bawah, atau yang paling spesifik terkait dengan suatu entity tidak dapat mengintervensi layer yang berposisi lebih atas, yang kemungkinan dapat menyebabkan kerusakan pada bisnis logic dari suatu usecase. 

- **Entities**
    > They encapsulate the most general and high-level rules. They are the least likely to change when something external changes.
    
    Dalam Hal ini diterapkan bahwa suatu object bertanggung jawab dalam mengatur behaviornya sendiri, dan tidak terpengaruh ketika terjadi perubahan dari sisi luar object tersebut. Hal ini diterapkan dengan tujuan agar suatu entity hanya bertanggung jawab untuk mengatur perilaku yang berkaitan dengan tujuan dari entity tersebut, dan tidak mempengaruhi layer yang lebih luar.
    
- **Use Cases**
    > These use cases orchestrate the flow of data to and from the entities, and direct those entities to use their enterprise wide business rules to achieve the goals of the use case.
    
    Dalam project ini diterapkan suatu layer yang bertugas sebagai orkestrator aliran data yang memungkinkan lebih dari satu entities terlibat didalamnya, hal ini diterapkan dengan tujuan agar ketika terjadi perubahan terkait bisnis proses yang berkaitan dengan alur data dari suatu usecase tertentu, maka dapat dilakukan dengan lebih mudah, karena ditempatkan pada satu tempat yang sama, dan tidak akan mempengaruhi layer entities.
    
    
- **Interface Adapters**
    > The software in this layer is a set of adapters that convert data from the format most convenient for the use cases and entities, to the format most convenient for some external agency such as the Database or the Web.
    
    Dalam Hal ini diterapkan untuk mengkonversi entity agar dapat di gunakan untuk external agency, dalam hal ini Database dan JSON representation, hal ini di terapkan dengan tujuan agar mempermudah dalam merepresentasikan suatu entity sesuai dengan kebutuhannya.

## How to Use

- **Build Docker-Compose file**

build docker-compose file, using this:

```
docker-compuse up
```
- **Create Database**

after the postgresql container is created, try to access the shell from postgresql, but firstly Switch to the Postgres user using this command:
    
```
su - postgres
```

and then run:

```
psql
```

after that create the database:

```
\c product_management
```

- **Setup The Configuration**

Make sure the config matches the database that was created previously

```
env: "development"
http_port: "9148"
postgresql:
    host: "localhost:5432"
    database: "product_management"
    username: "postgres"
    password: "root"
    sslmode: "disable"
    max_idle_conns: 2
    conn_max_lifetime: 3600000
    ping_interval: 5000
log_level: "debug"
```

- **Migrate The Table**

access the postgres shell, and run the table migration, that located in **db/create_product.sql**

```
CREATE TABLE IF NOT EXISTS "products" (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    description TEXT NOT NULL,
    quantity BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

- **Run The Project**

run the project using this commands:
```
go mod tidy
go run main.go server
```