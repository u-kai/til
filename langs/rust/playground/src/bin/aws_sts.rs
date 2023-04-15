#[tokio::main]
async fn main() {
    let config = aws_config::load_from_env().await;
    let client = aws_sdk_sts::Client::new(&config);
    let result = client
        .assume_role()
        .role_arn("arn:aws:iam::111815285043:role/delete_soon")
        .role_session_name("test_assume_role")
        .send()
        .await
        .unwrap();
    println!("{:#?}", result);
}
