# cards
Database for Anki Decks with reviewer functionality in the future

## todo
- [x] `server/graph/schema.resolvers.go` => handle file upload
- [ ] kerberos authentication
    - [x] admin view
    - [x] login popup
    - [x] jwt token for specific mutations
    - [ ] authentication against fs server 
- [x] `lib/modules/deck_selection_menu.dart` => file download: still not taking the new filename
- [x] search bar
- [x] `lib/modules/create_deck_dialog.dart` => language selection
- [x] `lib/modules/create_deck_dialog.dart` => field validation currently updates after a file was uploaded and not on text entry, which is bad => file upload currently has to be the last thing to do for the validation to work "properly"...
- [ ] localization
- [x] edit card view

## dev
1. get `mongodb` package and start the service
2. needed env var: `export JWT_SECRET_KEY=something`
3. in `server/` start the API endpoint via `go run server.go`
4. `flutter run -d web-server`

### frontend
This app is build upon the framework [Flutter](https://flutter.dev/). 
All frontend code is located inside the `frontend/lib/` folder and is structured in `modules` (big, outsourced parts of pages), `pages`, `utils` (helper functions) and `views` (reusable elements). All used constant values, as well as the queries and mutations to communicate to the backend can be found at `lib/constants.dart`.

### backend
The API is a [GraphQL](https://graphql.org/) instance, which is generated by [gqlgen](https://gqlgen.com/) and located in `server/`. The important files to look at are `server/graph/schema.graphqls` and `server/graph/schema.resolvers.go`. Former contains the datastructure to store metadata of deckfiles in a [MongoDB](https://www.mongodb.com/) database. The latter holds all logic to handle queries and mutations, requested by the frontend.
All uploaded deckfiles are stored within `server/deckfiles` and are automatically named by their sha256 hash like `<hash>.<apkg/colpkg>`.
The login fires against a LDAP/Kerberos instance and is handled in `server/utils/utils.go`.

## production
When publishing, two additional env vars are needed. In the future they are going to be placed inside a `.env` file.
```
ENVIRONMENT=production
LDAP_URL="ldaps://<ldap url>"
```