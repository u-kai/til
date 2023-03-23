//use actix_web::{web, App, HttpResponse, HttpServer, Responder};
use std::error::Error;

//#[actix_web::get("/")]
//async fn hello() -> impl Responder {
//HttpResponse::Ok().body("Hello world!")
//}

//#[actix_web::post("/echo")]
//async fn echo(req_body: String) -> impl Responder {
//HttpResponse::Ok().body(req_body)
//}

//async fn manual_hello() -> impl Responder {
//HttpResponse::Ok().body("Hey there!")
//}
//#[actix_web::main]
//async fn main() -> std::io::Result<()> {
//HttpServer::new(|| {
//App::new()
//.service(hello)
//.service(echo)
//.route("/hey", web::get().to(manual_hello))
//})
//.bind(("127.0.0.1", 8080))?
//.run()
//.await
//}
#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let url = "https://jsonplaceholder.typicode.com/posts/1";

    let client = reqwest::Client::new();

    let mut headers = reqwest::header::HeaderMap::new();
    headers.insert(
        reqwest::header::USER_AGENT,
        reqwest::header::HeaderValue::from_static("reqwest-tokio-example"),
    );

    let res = client
        .get(url)
        .headers(headers)
        .send()
        .await?
        .text()
        .await?;

    println!("{}", res);

    Ok(())
}
