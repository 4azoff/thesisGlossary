use actix_web::{get, App, HttpResponse, HttpServer, Responder};

#[get("/")]
async fn index() -> impl Responder {
    let terms = vec![
        "Programming language is a formal language used to form data structures and algorithms, that is, computational rules that can be executed by a computer.",
        "Rust is a general-purpose multi-paradigm compiled programming language that combines functional and procedural programming paradigms with a trait-based object system.",
        "Go is an open source compiled multi-threaded programming language from Google. It is considered a general-purpose language, but its main use is the development of web services and client-server applications.",
        "Server is a computer that stores data or performs certain service functions for other computers on the network.",
        "Client is a computer requesting some function or data from the server.",
        "Application is a task-oriented program designed for user interaction.",
        "«Client-server» is a computing or network architecture in which jobs or network load are distributed between service providers, called servers, and service customers, called clients.",
        "Backend is the internal part of the product, which is located on the server and is hidden from users."
    ];
    
    let mut content = String::new();
    content += "<html><head><link rel=\"stylesheet\" type=\"text/css\" href=\"/static/style.css\"></head><body><h1>Terms of work</h1><ul>";
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
