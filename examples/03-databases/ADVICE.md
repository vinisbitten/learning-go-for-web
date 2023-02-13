# Advice

For the examples **3,4 and 6** you will have to work with mysql. You can either install it on your computer or use a docker container. If you know how to use docker I recommend you to use it.

You can use the **dump.sql** file to create the database and the table you need for this example.

If you'll use mysql locally, you will have to create a database called **go_course** and a table called **posts**. The table should have the following structure:

| Field | Type         | Null | Key | Default | Extra          |
|-------|--------------|------|-----|---------|----------------|
| id    | int          | NO   | PRI | NULL    | auto_increment |
| title | varchar(255) | NO   |     | NULL    |                |
| body  | text         | NO   |     | NULL    |                |
