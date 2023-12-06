package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)




type Animals struct{
	Id uint16 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Birth uint16 `json:"birth"`
	Commands string `json:"commands"`

}
var posts=[]Animals{}
var showPost=Animals{}


	func main(){
		
		handleFunc()	


		// for  Flag == false {
		// 	fmt.Println("\n| 1-Показать базу данных \n"+"| 2-Добавить нового животного \n"+"| 3-Удалить животное \n"+"| 4-Закрыть меню\n")	
		// 	var number int
		// 	fmt.Print("Введите цифру: ")
		// 	fmt.Fscan(os.Stdin, &number) 
		// 	base(number)
		// }
        

	}
	func index(w http.ResponseWriter ,r *http.Request){
		t, err :=template.ParseFiles("templates/index.html","templates/header.html","templates/footer.html")
		if err!= nil{
			fmt.Fprintf(w,err.Error())
		}
		Db, err :=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/animals")
	    if err!=nil {
		   panic(err)
	    }
	    defer Db.Close()
		res, err := Db.Query("Select * From `animals`")
			if err!=nil {
				panic(err)
			}
			posts=[]Animals{}
			for res.Next(){
				var animal Animals
				err = res.Scan(&animal.Id, &animal.Name ,&animal.Type, &animal.Birth , &animal.Commands)
				if err!=nil {
					panic(err)
				}
				
				// a:=fmt.Sprintf("Name: %s \n Type: %s \n Age: %d \n Commands: %s \n", animal.Name, animal.Type, animal.Birth, animal.Commands)
				posts=append(posts,animal)
				
			}
		t.ExecuteTemplate(w,"index",posts)
	}
	func create(w http.ResponseWriter ,r *http.Request){
		t, err :=template.ParseFiles("templates/create.html","templates/header.html","templates/footer.html")
		if err!= nil{
			fmt.Fprintf(w,err.Error())
		}
		t.ExecuteTemplate(w,"create",nil)
	}
	func save_article( w http.ResponseWriter ,r *http.Request){
		
		nameAnimal:=r.FormValue("nameAnim")
		typeAnimal:=r.FormValue("typeAnim")
		dateBirth:=r.FormValue("birthAnim")
		commandsAnimal:=r.FormValue("commandsAnim")

		Db, err :=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/animals")
	    if err!=nil {
		   panic(err)
	    }
	    defer Db.Close()
		
		insert , err :=Db.Query("Insert Into `animals` (`name`,`type`,`birth`,`commands`) Values(?,?,?,?)",nameAnimal,typeAnimal, dateBirth,commandsAnimal)
		if err!=nil{
			panic(err)
		}
		defer insert.Close()
		http.Redirect(w,r,"/ ",http.StatusSeeOther)
	}
	func show_post( w http.ResponseWriter ,r *http.Request){
		t, err :=template.ParseFiles("templates/show.html","templates/header.html","templates/footer.html")
		if err!= nil{
			fmt.Fprintf(w,err.Error())
		}
		vars:=mux.Vars(r)
		Db, err :=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/animals")
	    if err!=nil {
		   panic(err)
	    }
	    defer Db.Close()
		//fmt.Fprintf(w, "ID: %v\n", vars["id"])
		res, err := Db.Query("Select * From `animals` where `id`=?",vars["id"])
			if err!=nil {
				panic(err)
			}
		showPost=Animals{}
		for res.Next(){
			var animal Animals
			err = res.Scan(&animal.Id, &animal.Name ,&animal.Type, &animal.Birth , &animal.Commands)
			if err!=nil {
				panic(err)
			}
				
			// a:=fmt.Sprintf("Name: %s \n Type: %s \n Age: %d \n Commands: %s \n", animal.Name, animal.Type, animal.Birth, animal.Commands)
			showPost=animal
				
			}
		t.ExecuteTemplate(w,"show",showPost)
	}



	func handleFunc(){
		rtr:=mux.NewRouter()
		rtr.HandleFunc("/", index).Methods("GET")
		rtr.HandleFunc("/create", create).Methods("GET")
		rtr.HandleFunc("/save_article", save_article).Methods("POST")
		rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")

		http.Handle("/",rtr)
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

		http.ListenAndServe(":8080",nil)
		

	} 
    
// 	func Base(num int)  {
		
// 		Db, err :=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/animals")
// 	   if err!=nil {
// 		   panic(err)
// 	   }
		
// 	   defer Db.Close()
//    }
			// switch num {
			// case 1:
			// 	open(Db)
			// case 2:
			// 	append(Db)
			// case 3:
			// 	delete(Db)
			// case 4:
			// 	Flag=true
			// }
		
	
	// func open(Db *sql.DB)  {
		
	        
			
	// 	// 
	// 	fmt.Println("Подключено к mysql")
	// }
	// func append(Db *sql.DB)  {
	// 	var nameAnimal string
	// 	var typeAnimal string
	// 	var dateBirth int
	// 	var commandsAnimal string
	// 	fmt.Print("Введите имя животного: ")
    // 	fmt.Fscan(os.Stdin, &nameAnimal)
	// 	fmt.Print("Введите тип животного(кошки,собаки,хомяки,лошади,верблюды,ослы): ")
    // 	fmt.Fscan(os.Stdin, &typeAnimal)
	// 	fmt.Print("Введите возраст животного: ")
    // 	fmt.Fscan(os.Stdin, &dateBirth)
	// 	fmt.Print("Введите комманду которую знает: ")
    // 	fmt.Fscan(os.Stdin, &commandsAnimal)
	// func Base()  {
	// 	Db, err :=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/animals")
	//     if err!=nil {
	// 	   panic(err)
	//     }
	//     defer Db.Close()
	// 	insert , err :=Db.Query("Insert Into `animals` (`name`,`type`,`birth`,`commands`) Values('Alex','vdcs',1,'ffffff')")
	// 	if err!=nil{
	// 		panic(err)
	// 	}
	// 	defer insert.Close()
	// }
	// }
	// func delete(Db *sql.DB){
	// 	var num int
	// 	fmt.Print("Введите id животного которого надо удалить : ")
    // 	fmt.Fscan(os.Stdin, &num)
	// 	insert , err :=Db.Query("Delete  from `animals` where id=? ",num)
	// 	if err!=nil{
	// 		panic(err)
	// 	}
	// 	defer insert.Close()
	// }
