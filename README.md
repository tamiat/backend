

# backend
Tamiat CMS backend

#### Development instructions
- You should rename .env.example with .env and put all environment values
  ex: HOST=localhost.
- If you want to add a new type of errors, You can add it in [errs package](pkg/errs/errors.go).

#### Running instructions

- To test  the endpoints, import the [collection]() in [Postman](https://www.postman.com/), you can check how to import it from [here](https://kb.datamotion.com/?ht_kb=postman-instructions-for-exporting-and-importing).
- After every  `pull` you can run `soda migrate` to update your local database structure if there are any updates.
#### Database on local
1. Create postgresql database  from terminal by doing:

    ```
     sudo -i -u postgres
     ```  
   ```  
   psql  
   ```  
   ```  
   CREATE DATABSE db_name;  
    ```
2. Add datasource to goland (optional, you can use pgadmin4 or dbeaver or you can use postgres from command line)  
   leave all settings and just add the postgres username as shown in the picture.  
   ![1](https://user-images.githubusercontent.com/49435053/132143481-3b7f28da-55da-4d48-adca-affa7afb02b8.png)

3. Environment variables:
   There are 5 environment variables related to database connection</br>

- PASS=< value_for_password_of_postgres_database >

- HOST=localhost </br>

- DBNAME=<db_name> ;  // the same name of database that was created in postgres as shown above</br>

- DBPORT=5432 ; </br>

- PORT=8080 ; </br>

\
4. Install soda migration tool:
1. In linux:
- open .profile
  ```  
  nano .profile  
  ```  
- add this at the end of the file: <\br>
  ```
  export PATH=$HOME/go/bin:$PATH  
  ```
  ```  
  source .profile  
  ```  
2. In mac:  
   same instructions as linux but open .zprofile

3. In windows:  
   add this environment variable:
   ```  
   C:\Users\<your_username>\go\bin  
   ```  
- then in the working directory of the project:
  ```
  soda migrate
  ```
- for more info about soda migration and how it works you can check this [link](https://gobuffalo.io/en/docs/db/toolbox)

#### Send verification code feature
We use sendgrid api, so to test this feature you have to:
- create an account in [sendgrid](https://sendgrid.com/)
-   create api key
-  create a template and replace TEMPLATE_ID env variable with your own template id.
#### Swagger
Use this url to see swagger page
```  
http://localhost:8080/swagger/index.html#/  
```