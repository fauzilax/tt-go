# tt-go


<h2>How to run local</h2>

- Clone from this repository

  ``` $git clone https://github.com/fauzilax/gdt-api.git ```

- change directory 

  ``` $cd gdt-api```
 
- Run The Program

  ``` $go run main.go ```


<h3> All API in this repo</h3>

- MEMBERS TABLE

  ```POST /insert``` <br>
  ```PUT /update/:id_member ```<br>
  ```DELETE /delete/:id_member ```<br>
  ```GET /selectallmember ```<br>

- REVIEW_PRODUCTS TABLE

  ```GET /selectidproduct/:id_product``` <br>

- LIKE_REVIEWS TABLE

    insert <br>
    ```POST /like ```<br>
    delete<br>
    ```DELETE /dislike``` <br>
