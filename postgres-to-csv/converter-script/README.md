### converter-script: to convert postgres data into csv file

* clone this respository then run this command on your mac terminal
    ```bash
    # host=your_hostname port=your_portno. dbname=your_db_name user=your_username password=your_password  
    psql "host=localhost port=5432 dbname=college user=varun password=something" -af my_query.sql 
     ```