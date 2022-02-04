# restAlbum

Basic REST API written in GO using Gin Web Framework.

Gin Documentation: https://pkg.go.dev/github.com/gin-gonic/gin

3 Basic calls:

Get - Retrieves all albums, no arguments or JSON.

Get By Id - Retrieves an album by Id, Id is sent via a command line argument. No JSON.

Post Album - Creates a new album using request supplied JSON.

Commands to practice requests from the command prompt on windows:
curl http://localhost:8080/albums - returns all albums
curl http://localhost:8080/albums/{id} - returns 1 album with a matching id
curl -X POST -H "Content-type: application/json" --data "{\"id\": \"{id}\", \"title\": \"{album-name}\",\"artist\": \"{artist-name}\", \"price\": {price}}" http://localhost:8080/albums
curl -X "DELETE" http://localhost:8080/albums/{id} - removes album with matching id and returns remaining albums

For examples above replace curly braces and words with variables for example:
{id} replace with 4 or 5 or 6
{album-name} replace with Hey Ya!
{artist-name} replace with OutKast
{price} replace with 14.99 or 20.99
