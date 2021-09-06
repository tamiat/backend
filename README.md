
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
- we will add 6 environment variables using the following format: </br>

PASS=< value for password> ; export PASS </br>

HOST=localhost ; export HOST </br>

DBNAME=cms ; export DBNAME // the same name of database that was created in postgres </br>

DBPORT=5432 ; export DBPORT </br>

PORT=8080 ; export PORT </br>

SECRET=< value for jwt secret > ; export SECRET </br>


then run this command in the project root directory:
```
source ~/.bashrc
```
4. Install soda migration tool:
- In linux:
```
cd ..
```
```
nano .profile
```
add this at the end of the file: <\br>

```
export PATH=$HOME/go/bin:$PATH
```
```
source .profile
```
- In mac:
same instructions as linux but open .zprofile

- In windows:
follow the instructions in this link https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/ to add new env variable, and add this:
```
C:\Users\<your_username>\go\bin
```
then in the working directory of the project:
```
cd pkg
```
```
soda migrate
```
