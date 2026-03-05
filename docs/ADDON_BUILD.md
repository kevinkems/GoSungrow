# Home Assistant Addon – Build-Fehler beheben (404)

## Problem

Beim Bau des Home-Assistant-Addons erscheint:

```text
HTTP request sent, awaiting response... 404 Not Found
https://api.github.com/repos/kevinkems/GoSungrow/releases/latest
```

**Ursache:** Das Addon lädt die GoSungrow-Binary beim Docker-Build von den GitHub-Releases. Solange unter **kevinkems/GoSungrow** kein Release existiert, liefert `/releases/latest` 404.

## Lösung: Release erstellen

Sobald unter **kevinkems/GoSungrow** mindestens ein Release mit den passenden Assets (z. B. `GoSungrow-linux_arm64.tar.gz`) existiert, funktioniert der Addon-Build.

### 1. Version prüfen

In `defaults/const.go` steht die aktuelle Version (z. B. `BinaryVersion = "3.0.8"`). Das Release-Tag muss dazu passen (z. B. `v3.0.8`).

### 2. Release mit Goreleaser erstellen

```bash
# Im Projektroot (GoSungrow)
./bin/release.sh
```

Das Skript:

- prüft die Version aus `defaults/const.go`
- erstellt einen Git-Tag (z. B. `v3.0.8`)
- pusht und startet `goreleaser release --rm-dist`
- erzeugt die gewohnten Archive (u. a. `GoSungrow-linux_arm64.tar.gz`) und hängt sie am GitHub-Release an

### 3. Alternativ: manuell

```bash
# Version anpassen, z. B. in defaults/const.go: BinaryVersion = "3.0.8"
VERSION="v3.0.8"

git add .
git commit -m "Release $VERSION"
git tag "$VERSION"
git push && git push --tags

goreleaser release --rm-dist
```

### 4. Addon erneut bauen

Nach dem Erstellen des Releases den Addon-Build in Home Assistant erneut ausführen. Der Aufruf von `https://api.github.com/repos/kevinkems/GoSungrow/releases/latest` liefert dann ein gültiges Release mit den benötigten Assets.

## Hinweis

Das Addon (Dockerfile, `setup.sh`, `getRelease.sh`) stammt typischerweise aus einem Addon-Repo (z. B. Fork von [MickMake/HomeAssistantAddons](https://github.com/MickMake/HomeAssistantAddons)). Dort wird die Repo-URL für die Releases konfiguriert (z. B. `kevinkems/GoSungrow`). Solange diese URL auf ein Repo **mit mindestens einem Release** zeigt, baut das Addon erfolgreich.
