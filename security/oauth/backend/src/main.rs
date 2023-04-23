use std::collections::HashMap;

use actix_web::{web, App, HttpResponse, HttpServer, Responder};
use reqwest::{
    header::{ACCEPT, AUTHORIZATION, CONTENT_TYPE, USER_AGENT},
    Client,
};
use serde_json::{json, Value};

#[actix_web::get("/")]
async fn hello() -> impl Responder {
    HttpResponse::Ok().body("Hello world!")
}

#[actix_web::post("/echo")]
async fn echo(req_body: String) -> impl Responder {
    HttpResponse::Ok().body(req_body)
}

#[actix_web::get("/callback")]
async fn callback(query: web::Query<HashMap<String, String>>) -> impl Responder {
    println!("{:#?}", query);
    let body = json!(
    {
            "client_id": std::env::var("GITHUB_OAUTH_CLIENT_ID").unwrap(),
            "client_secret":std::env::var("GITHUB_OAUTH_CLIENT_SECRET").unwrap(),
            "code":query.get("code").unwrap(),
            //"state":query.get("state").unwrap()
    });
    println!("body {:#?}", body.to_string());
    let response = Client::new()
        .post("https://github.com/login/oauth/access_token")
        .header(CONTENT_TYPE, "application/json")
        .header(ACCEPT, "application/json")
        .body(body.to_string())
        .send()
        .await
        .unwrap();

    let res = response.text().await.unwrap();
    println!("{:#?}", res);
    let token: Res = serde_json::from_str(&res).unwrap();
    println!("{:#?}", token);
    let response = Client::new()
        .get("https://api.github.com/user/repos")
        .header(CONTENT_TYPE, "application/json")
        .header(ACCEPT, "application/json")
        .header(USER_AGENT, "rust-program")
        .header("Authorization", format!("Bearer {}", token.access_token));

    println!("{:#?}", response);

    let response = response.send().await.unwrap();
    let header = response.headers();
    println!("header {:#?}", header);
    let response = response.text().await.unwrap();
    let array = serde_json::from_str::<Value>(&response);
    match array {
        Ok(Value::Array(a)) => HttpResponse::Ok().body(a.len().to_string()),
        _ => HttpResponse::Ok().body(response),
    }
}

#[derive(serde::Deserialize, Debug)]
struct Res {
    access_token: String,
}
#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| App::new().service(hello).service(echo).service(callback))
        .bind(("127.0.0.1", 8000))?
        .run()
        .await
}
