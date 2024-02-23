# cards
Database for Anki Decks with reviewer functionality in the future

## dev ([todos](https://github.com/FachschaftMathPhysInfo/cards/issues/1))
1. `docker compose build`
2. `docker compose up -d && docker compose logs -f` (frontend: `localhost:8081`, backend: `localhost:8080`)
3. (debugging the frontend: in `frontend/` execute `flutter run -d web-server`)

### frontend
This app is build upon the framework [Flutter](https://flutter.dev/). 
All frontend code is located inside the `frontend/lib/` folder and is structured in `modules` (big, outsourced parts of pages), `pages`, `utils` (helper functions) and `views` (reusable elements). All used constant values, as well as the queries and mutations to communicate to the backend can be found at `lib/constants.dart`.

### backend
The API is a [GraphQL](https://graphql.org/) instance, which is generated by [gqlgen](https://gqlgen.com/) and located in `server/`. The important files to look at are `server/graph/schema.graphqls` and `server/graph/schema.resolvers.go`. Former contains the datastructure to store metadata of deckfiles in a [MongoDB](https://www.mongodb.com/) database. The latter holds all logic to handle queries and mutations, requested by the frontend.
All uploaded deckfiles are stored within `server/deckfiles` and are automatically named by their sha256 hash like `<hash>.<apkg/colpkg>`.
The login fires against a LDAP/Kerberos instance and is handled in `server/utils/utils.go`.
