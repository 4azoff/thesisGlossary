use actix_web::{get, App, HttpResponse, HttpServer, Responder};

#[get("/")]
async fn index() -> impl Responder {
    let terms = vec![
        "Rust",
        "Actix-web",
        "HTML",
        "Web development",
    ];
    
    let mut content = String::new();
    content += "<html><head><link rel=\"stylesheet\" type=\"text/css\" href=\"/static/style.css\"></head><body><h1>Terms:</h1><ul>";
    for term in terms {
        content += &format!("<li><a href=\"#\">{}</a></li>", term);
    }
    content += "</ul></body></html>";
    
    HttpResponse::Ok().body(content)
}

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(actix_files::Files::new("/static", "./static").show_files_listing())
            .service(index)
            
    })
    .bind(("0.0.0.0", 8088))?
    .run()
    .await
}
