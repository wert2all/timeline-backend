use axum::routing::post;
use axum::Router;
use diesel_async::pooled_connection::deadpool::Pool;
use diesel_async::pooled_connection::AsyncDieselConnectionManager;
use shuttle_axum::ShuttleAxum;

mod db;
mod graphql;
mod state;

#[shuttle_runtime::main]
async fn main(#[shuttle_shared_db::Postgres] conn_str: String) -> ShuttleAxum {
    let connection = &mut db::connect(conn_str.clone());
    db::run_migration(connection);
    let config =
        AsyncDieselConnectionManager::<diesel_async::AsyncPgConnection>::new(conn_str.clone());
    let pool: Pool<AsyncDieselConnectionManager<diesel_async:AsyncPgConnection>> =
        Pool::builder(config).build().unwrap();
    let _connection = pool.get().await.unwrap();
    let state = state::AppState::new();

    let router = Router::new()
        .route("/graphql", post(graphql::handler))
        .with_state(state);

    Ok(router.into())
}
