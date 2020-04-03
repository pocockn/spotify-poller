# spotify-poller

Polls the Spotify-API for new tracks added to a playlist.

## Getting Started

```
make run SPOTIFY_CLIENT_SECRET="XXX"
```

### Installing

```
make install
```

## Running the tests

```
make test
```


## Built With

* [Go](https://golang.org/)
* [Spotify API Go Wrapper](https://github.com/zmb3/spotify)
* [GORM](https://gorm.io/) - ORM Library

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/pocockn/recs-api/tags). 

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## TODO

- Allow for different playlists to be polled
- CI
- Pass in Client ID through ENV var
