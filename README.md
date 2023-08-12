<div align="center">
  <h1>InCrowd Backend</h1>
  <blockquote>System that processes data from external feed providers and provides and API to requests that data</blockquote>
</div>

<br/>

## ⚙️ Usage and examples

### Run project

```
make docker-up  //Cleans up containers and then starts up InCrowd Backend and its dependecies
```

### Request examples

```
localhost:8080/monitor_probe
localhost:8080/provider/realise/v1/teams/t94/news/611106
localhost:8080/provider/realise/v1/teams/t94/news?page=0&count=20&sort=published&order=desc
```


## 📜 Information

### Mongo Express
After running ```make docker-up```, it also starts up Mongo Express, so you can easily check mongoDB data. The service is running on ```http://localhost:8081```. 

### Pagination and sorting
The API allows adding pagination and sorting to list the articles. Provided parameters: page, count, sort (published, id), order (asc, desc).

### Project Architecture
A quick Medium reading about Hexagonal Architecture in Go: [Hexagonal Architecture](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3)

### TODO
- Add testing
- Add swagger page
- Expose metrics and run Grafana on Docker container
- Retrieve articles details concurrently
- Handle metadata for pagination and sorting in a more generic way, for example, using a middleware to store it within the context of the request
- Etc...
