package main

import (
  "fmt"
  "net"
  "net/http"
  "net/http/fcgi"
  "code.google.com/p/gorest"
)

func main() {
  fmt.Println("Hello!")
  listener, err := net.Listen("tcp", ":8080")
  if (err != nil) {
    fmt.Println("Error!")
    fmt.Println(err)
  }
  srv := http.NewServeMux()
  srv.HandleFunc("/", gorest.handle())
  fcgi.Serve(listener, srv)
}

type accessToken struct{
  
}

//************************Define Service***************************

type AuthService struct{
  gorest.RestService    `root:"/" consumes:"application/json" produces:"application/json"`

  authenticate  gorest.EndPoint `method:"POST" path:"/authenticate/" postdata:"authUser" output:"accessToken"`
}

type PokerService struct{
  //Service level config
  gorest.RestService    `root:"/poker/" consumes:"application/json" produces:"application/json"`

  //End-Point level configs: Field names must be the same as the corresponding method names,
  // but not-exported (starts with lowercase)

  tableList     gorest.EndPoint `method:"GET" path:"/table/" output:"[]Table"`
  tableDetails  gorest.EndPoint `method:"GET" path:"/table/{Id:int}" output:"Table"`
  playerList    gorest.EndPoint `method:"GET" path:"/player/" output:"[]Player"`
  playerDetails gorest.EndPoint `method:"GET" path:"/player/{Id:int}" output:"Player"`
  //playerStats   gorest.EndPoint `method:"GET" path:"/player/{Id:int}/" output:"PlayerStats"`
}

type UserService struct{
  gorest.RestService    `root:"/user/" consumes:"application/json" produces:"application/json"`

  userDetails   gorest.EndPoint `method:"GET" path:"/{Id:int}" output:"User"`
  userCreate    gorest.EndPoint `method:"POST" path:"/" postdata:"User"`
}

//Handler Methods: Method names must be the same as in config, but exported (starts with uppercase)

func(serv PokerService) TableDetails(Id int) (t Table){
    if table,found:=tableStore[Id];found{
        t =table
        return
    }
    serv.ResponseBuilder().SetResponseCode(404).Overide(true)  //Overide causes the entity returned by the method to be ignored. Other wise it would send back zeroed object
    return
}

func(serv PokerService) TableList()[]Table{
    serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
    return tableStore
}

func(serv PokerService) PlayerDetails(Id int) (p Player){
    if player,found:=playerStore[Id];found{
        p =player
        return
    }
    serv.ResponseBuilder().SetResponseCode(404).Overide(true)  //Overide causes the entity returned by the method to be ignored. Other wise it would send back zeroed object
    return
}

func(serv PokerService) PlayerList()[]Player{
    serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
    return playerStore
}

// func(serv OrderService) AddItem(i Item){

//     for _,item:=range itemStore{
//         if item.Id == i.Id{
//             item=i
//             serv.ResponseBuilder().SetResponseCode(200) //Updated http 200, or you could just return without setting this. 200 is the default for POST
//             return
//         }
//     }

//     //Item Id not in database, so create new
//     i.Id = len(itemStore)
//     itemStore=append(itemStore,i)

//     serv.ResponseBuilder().Created("http://localhost:8787/orders-service/items/"+string(i.Id)) //Created, http 201
// }

// //On the method parameters, the posted data(http-entity) is always first, followed by the URL mapped parameters
// func(serv OrderService) PlaceOrder(order Order,UserId int,AskForDiscount bool){
//     order.Id = len(orderStore)

//     if user,found:= userStore[UserId];found{
//           if item,exists:=findItem(order.ItemId);exists{
//                 itemStore[item.Id].AvailableStock--

//                 if AskForDiscount && order.Amount >5{
//                     order.Discount = 2.5
//                 }
//                 order.Id=len(orderStore)
//                 order.UserId=UserId
//                 order.Cancelled=false
//                 orderStore=append(orderStore,order)
//                 user.OrderIds=append(user.OrderIds,order.Id)

//                 userStore[user.Id]=user

//                 serv.ResponseBuilder().SetResponseCode(201).Location("http://localhost:8787/orders-service/orders/"+string(order.Id))//Created
//                 return

//           } else{
//               serv.ResponseBuilder().SetResponseCode(404).WriteAndOveride([]byte("Item not found"))//You can still manually place an entity on the response, even on a POST
//               return
//           }
//     }

//     serv.ResponseBuilder().SetResponseCode(404).WriteAndOveride([]byte("User not found"))
//     return
// }
// func(serv OrderService) ViewOrder(id int) (retOrder Order){
//      for _,order:=range orderStore{
//         if order.Id == id{
//             retOrder = order
//             return
//         }
//      }
//      serv.ResponseBuilder().SetResponseCode(404).Overide(true)
//      return
// }
// func(serv OrderService) DeleteOrder(id int) {
//      for pos,order:=range orderStore{
//         if order.Id == id{
//             order.Cancelled =true
//             orderStore[pos]=order
//             return               //Default http code for DELETE is 200
//         }
//      }
//      serv.ResponseBuilder().SetResponseCode(404).Overide(true)
//      return
// }