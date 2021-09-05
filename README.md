
# backend
Tamiat CMS backend

#### Development instructions

- If you want to add a new type of errors, You can add it in [errs package](pkg/errs/errors.go).

#### Running instructions

- You should export environment variables from terminal 
  before running main.go
  ex: export HOST=localhost. 
- To test  the endpoints, import the [collection](postman%20collection/backend.postman_collection.json) in [Postman](https://www.postman.com/) and you can check how to import it from [here](https://kb.datamotion.com/?ht_kb=postman-instructions-for-exporting-and-importing).
#### Database on local
1. Create postgresql database  from terminal by doing:
 ```
 sudo -i -u postgres
```
```
psql
```
```
CREATE DATABSE cms;
 ```
 2. Add datasource to goland:
 leave all settings and just add the postgres username as shown in the picture.
 ![1](https://user-images.githubusercontent.com/49435053/132143481-3b7f28da-55da-4d48-adca-affa7afb02b8.png)

 3. Environment variables:
 - open .bashrc file.
- we will add 6 environment variables using the following format:
PASS=< value for password> ; export PASS
HOST=localhost ; export HOST
DBNAME=cms ; export DBNAME // the same name of database that was created in postgres
DBPORT=5432 ; export DBPORT
PORT=8080 ; export PORT
SECRET=< value for jwt secret > ; export SECRET

then run this command in the project root directory:
```
source ~/.bashrc
```
4. Install soda migration tool:
use these commands in the project root directory to install soda:
```
go get -u -v -tags sqlite github.com/gobuffalo/pop/...
```
```
go install -tags sqlite github.com/gobuffalo/pop/soda
```
then 
```
cd pkg
```
```
soda migrate
```
