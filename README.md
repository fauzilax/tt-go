# tt-go


<h2>How to run local</h2>

- Clone from this repository

  ``` $git clone https://github.com/fauzilax/gdt-api.git ```

- change directory 

  ``` $cd gdt-api```
 
- Run The Program

  ``` $go run main.go ```


<h3> all API in this repo</h3>

- MEMBERS TABLE

  a. POST /insert <br>
  b. PUT /update/:id_member <br>
  c. DELETE /delete/:id_member <br>
  d. GET /selectallmember <br>

- REVIEW_PRODUCTS TABLE

  e. GET /selectidproduct/:id_product <br>

- LIKE_REVIEWS TABLE

  f. insert & delete<br>
      POST "/like" <br>
      DELETE "/dislike" <br>
