# sushi-roll-api
Repo ini adalah sebuah rest api sederhana dengan menggunakan `gorilla mux` tanpa menggunakan database. Jadi ketika server dihentikan maka semua data yang telah di-create akan hilang.  
untuk menjalankan ketik `go run main.go`.
# body request postman untuk create dan update
```
 {  
  "name": "insert name",  
  "description": "insert description",  
  "ingredients": "insert ingredients"  
 }  
``` 

 # test postman
Berikut ini adalah HTTP method beserta route-nya  
`localhost:5000/`  
  
GET "/sushi" &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;//get all sushi  
GET "/sushi/{id}" &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;  //get sushi by id  
POST "/sushi" &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; //create sushi  
PUT "/sushi/{id}" &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;   //update sushi by id  
DELETE "/sushi/{id}" &nbsp;&nbsp;  //delete sushi by id  

