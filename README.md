

# backend
Tamiat CMS backend

### Instructions applied on linux os
#### Prerequisites
- [Golang 1.6 or higher](https://go.dev/doc/install)
  You may use these steps to install go on your Linux machine or use the previous link to get the latest versions of go.
    1. Open your terminal and make sure you are at home path
    2. Download the go package.
       ```wget https://go.dev/dl/go1.17.7.linux-amd64.tar.gz```
    3. Extract go files
       ```sudo tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz```
    4. Add the following line to /home/<your_username/.profile file by using your preferred text editor.
       ```export PATH=$PATH:/usr/local/go/bin```
    5. To apply the changes made to .proile file run:
       ```source $HOME/.profile```
    6. Make sure you installed go by running:
       ```go version```
       You will get output message like this:
       ```go version go1.17.7 linux/amd64```
- Postgresql
  Follow the instructions here: (https://www.postgresql.org/download/)

#### Setting up the project
1. Clone the repository at any path
   ```git clone https://github.com/tamiat/backend.git```
2. ```cd backend```
3. Create postgresql database and set password for postgres user
    - ```sudo -i -u postgres```
    - ```psql```
    - ``` CREATE DATABASE <database_name>;  ```
    - ``` ALTER USER postgres WITH PASSWORD '<new_password>';```
4. type `exit` twice to return to the terminal.
5. Rename .env.example file to .env
   ```mv .env.example .env```
   This file contains all necessarily environment variables.


| Environment Variable  | Explanation |
| ------------- | ------------- |
| HOST  | ex: localhost  |
| DBNAME  | database name  |
| DBPASS  | database password  |
| DBPORT  | database port, set to 5432  |
| PORT  | which port number to use for endpoints, ex: 8080  |
| DBUSER  | database user, ex: postgres  |
| JWT_SECRET | string to set secret of jwt  |
| EMAIL_SENDER  | the email that will send confirmation emails to users. It should be activated in sendgrid profile   |
| TEMPLATE_ID  | which template id sendgrid uses as a format to send emails  |
| SENDGRID_API_KEY  | api key associated with the sender account |

##### Sendgrid is a thirdparty we use to send emails
6. Set environment variables using your preferred text editor
7. Back to backend directory, run:
   ```go get github.com/gobuffalo/pop/...```
8. Edit ~/.profile and add this line into it.
   ```export PATH=$PATH:~/go/bin```
9. Run the following command to apply changes to .profile
   ```source ~/.profile```
10. Run the following command
    ```which soda```
    It should give output like this:
    ```/home/<user_name>/go/bin/soda ```
11. Now run the following to create all database tables
    ```soda migrate up```

#### Running instructions
In the project root directory run
```go run main.go```


