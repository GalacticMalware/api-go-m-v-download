# 🎬🎵 Video & Music Downloader API

API REST escrita **100% en Go nativo** (sin frameworks) para descargar videos de YouTube en 720p/1080p y extraer audio en MP3/MP4.

Usa `yt-dlp` y `ffmpeg` como herramientas externas (invocadas vía `os/exec`), pero todo el servidor HTTP, enrutado, JSON y logging es Go puro (`net/http`, `log/slog`).

---

## 📋 Tabla de contenidos

- [Características](#-características)
- [Arquitectura](#-arquitectura)
- [Estructura del proyecto](#-estructura-del-proyecto)
- [Requisitos previos](#-requisitos-previos)
- [Instalación](#-instalación)
  - [Windows](#-windows)
  - [Linux](#-linux)
  - [macOS](#-macos)
- [Verificación e instalación de rutas (Paths)](#-verificación-e-instalación-de-rutas-paths)
- [Configuración](#-configuración)
- [Ejecución](#-ejecución)
- [Endpoints de la API](#-endpoints-de-la-api)
- [Ejemplos con curl](#-ejemplos-con-curl)
- [Ejecución como servicio (systemd)](#-ejecución-como-servicio-systemd)
- [Mantenimiento](#-mantenimiento)
- [Solución de problemas](#-solución-de-problemas)

---

## ✨ Características

- **Descarga de video** en `1080p` y `720p` (con audio mergeado vía ffmpeg).
- **Descarga de audio** en `mp3` y `mp4`/`m4a`.
- **Validación de campos requeridos** (`url` y `format`).
- **Validación de formatos** (rechaza formatos no permitidos con `Formato incorrecto`).
- **Logging estructurado** en consola y archivo (`logs/app.log`).
- **Middleware** de logging HTTP y recuperación de pánicos.
- **Arquitectura hexagonal** (puertos y adaptadores) → fácil de cambiar el motor de descarga.
- **Shutdown graceful**.

---

## 🏗 Arquitectura

Se utiliza **Clean Architecture / Arquitectura Hexagonal**:

```
HTTP Request
     │
     ▼
[ middleware: Logging / Recover ]
     │
     ▼
[ router: POST /api/v1/download/videos|music ]
     │
     ▼
[ handler ] ──(JSON)──► domain.Request
     │
     ▼
[ usecase ]  ◄── valida campos y formatos
     │
     ▼
[ ports.Downloader ] ◄── interfaz
     │
     ▼
[ adapter/ytdlp ] ──► exec yt-dlp + ffmpeg
     │
     ▼
domain.DownloadResult → JSON al cliente
```

---

## 📁 Estructura del proyecto

```
video-music-downloader/
├── cmd/
│   └── api/
│       └── main.go                      # Punto de entrada
├── internal/
│   ├── config/config.go                 # Configuración (env vars)
│   ├── domain/                          # Entidades y errores de negocio
│   │   ├── download.go
│   │   └── errors.go
│   ├── usecase/                         # Casos de uso
│   │   ├── video_usecase.go
│   │   └── music_usecase.go
│   ├── ports/downloader.go              # Interfaces (contratos)
│   ├── adapter/ytdlp/ytdlp_adapter.go   # Implementación con yt-dlp
│   ├── handler/                         # Capa HTTP
│   │   ├── video_handler.go
│   │   ├── music_handler.go
│   │   └── response.go
│   ├── middleware/                      # Middlewares
│   │   ├── logging.go
│   │   └── recover.go
│   └── router/router.go                 # Enrutador net/http
├── pkg/logger/logger.go                 # Logger (log/slog)
├── downloads/                           # Archivos descargados
├── logs/                                # Logs de la app
├── .env.example
├── .gitignore
├── go.mod
└── README.md
```

---

## ⚙️ Requisitos previos

| Herramienta            | Versión mínima     | Uso                                                           |
| ---------------------- | ------------------ | ------------------------------------------------------------- |
| **Go**                 | 1.22+              | Compilar/ejecutar la API (`ServeMux` con métodos)             |
| **yt-dlp**             | Última             | Motor de descarga                                             |
| **ffmpeg**             | Última             | Merge de audio+video y conversión a MP3                       |
| **ffprobe**            | (viene con ffmpeg) | Post-procesado de audio                                       |
| **Node.js** o **Deno** | Última             | JS runtime para YouTube (evita warnings y formatos faltantes) |

> ⚠️ **ffmpeg incluye ffprobe**. Asegúrate de que **ambos** estén disponibles.

---

## 🚀 Instalación

### 🪟 Windows

#### 1. Go

```powershell
winget install GoLang.Go
```

O descarga desde https://go.dev/dl/ y añade `C:\Program Files\Go\bin` al PATH.

**Verificar:**

```powershell
go version
```

#### 2. yt-dlp

```powershell
winget install yt-dlp
```

#### 3. ffmpeg

```powershell
winget install Gyan.FFmpeg
```

#### 4. JS runtime (Node.js o Deno)

```powershell
# Opción A: Node.js
winget install OpenJS.NodeJS

# Opción B: Deno (el que yt-dlp usa por defecto)
winget install DenoLand.Deno
```

**Cierra y reabre PowerShell** para que tome las variables de entorno.

#### 5. Verificar e instalar rutas del PATH

```powershell
where.exe yt-dlp
where.exe ffmpeg
where.exe ffprobe
where.exe node
```

**Salida esperada (ejemplo):**

```
C:\Users\<user>\AppData\Local\Microsoft\WinGet\Links\yt-dlp.exe
C:\Users\<user>\AppData\Local\Microsoft\WinGet\Links\ffmpeg.exe
C:\Users\<user>\AppData\Local\Microsoft\WinGet\Links\ffprobe.exe
C:\Program Files\nodejs\node.exe
```

Si algún comando da **"no se reconoce"**, usa la ruta absoluta del `.exe` en tu `.env` (ver sección [Configuración](#-configuración)).

**Valores para `.env` en Windows (usando barras `/` o dobles `\\`):**

```env
YTDLP_PATH=C:/Users/<user>/AppData/Local/Microsoft/WinGet/Links/yt-dlp.exe
FFMPEG_PATH=C:/Users/<user>/AppData/Local/Microsoft/WinGet/Links
```

> ⚠️ En `.env` en Windows **evita backslashes sueltos** (`C:\ffmpeg`) porque se interpretan como escapes. Usa `C:/ffmpeg` o `C:\\ffmpeg`.

---

### 🐧 Linux

#### 1. Go (Ubuntu/Debian)

```bash
GO_VERSION="1.22.5"
wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

> Para ARM: usa `linux-arm64`.

#### 2. ffmpeg

```bash
# Ubuntu/Debian
sudo apt update && sudo apt install -y ffmpeg

# RHEL/Fedora
sudo dnf install -y ffmpeg

# Arch
sudo pacman -S ffmpeg
```

#### 3. yt-dlp

```bash
sudo wget https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -O /usr/local/bin/yt-dlp
sudo chmod a+rx /usr/local/bin/yt-dlp
```

#### 4. JS runtime

```bash
# Node.js (Debian/Ubuntu)
curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
sudo apt install -y nodejs

# Deno
curl -fsSL https://deno.land/install.sh | sh
```

#### 5. Verificar rutas

```bash
which yt-dlp
which ffmpeg
which ffprobe
which node
```

**Salida esperada:**

```
/usr/local/bin/yt-dlp
/usr/bin/ffmpeg
/usr/bin/ffprobe
/usr/bin/node
```

**Valores para `.env` en Linux:**

```env
YTDLP_PATH=/usr/local/bin/yt-dlp
FFMPEG_PATH=/usr/bin
```

---

### 🍎 macOS

#### 1. Homebrew (si no lo tienes)

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

#### 2. Instalar todo con Homebrew

```bash
brew install go yt-dlp ffmpeg node
```

> Puedes usar `deno` en lugar de `node` si prefieres: `brew install deno`.

#### 3. Verificar rutas

```bash
which yt-dlp
which ffmpeg
which ffprobe
which node
```

**Salida esperada (según arquitectura):**

```
/opt/homebrew/bin/yt-dlp     # Apple Silicon (M1/M2/M3)
/opt/homebrew/bin/ffmpeg
/opt/homebrew/bin/ffprobe
/opt/homebrew/bin/node
```

**Valores para `.env` en macOS:**

```env
YTDLP_PATH=/opt/homebrew/bin/yt-dlp
FFMPEG_PATH=/opt/homebrew/bin
```

> En Intel Mac suele ser `/usr/local/bin` en vez de `/opt/homebrew/bin`.

---

## 🔍 Verificación e instalación de rutas (Paths)

Ejecuta **todos** estos comandos. Si alguno falla, ese componente no está en el PATH y hay que usar su ruta absoluta.

| Comando de verificación               | Qué confirma       |
| ------------------------------------- | ------------------ |
| `go version`                          | Go instalado       |
| `yt-dlp --version`                    | yt-dlp en PATH     |
| `ffmpeg -version`                     | ffmpeg en PATH     |
| `ffprobe -version`                    | ffprobe en PATH    |
| `node --version` (o `deno --version`) | JS runtime en PATH |

**Ubicar cada binario y su carpeta:**

| SO          | Comando para ubicar | Ejemplo de salida                |
| ----------- | ------------------- | -------------------------------- |
| Windows     | `where.exe yt-dlp`  | `C:\...\WinGet\Links\yt-dlp.exe` |
| Windows     | `where.exe ffmpeg`  | `C:\...\WinGet\Links\ffmpeg.exe` |
| Linux/macOS | `which yt-dlp`      | `/usr/local/bin/yt-dlp`          |
| Linux/macOS | `which ffmpeg`      | `/usr/bin/ffmpeg`                |

**Nombre de las variables en el `.env`:**

| Variable       | Qué debe apuntar                                    | Ejemplo Windows       | Ejemplo Linux/macOS                |
| -------------- | --------------------------------------------------- | --------------------- | ---------------------------------- |
| `YTDLP_PATH`   | Ruta al **binario** yt-dlp                          | `C:/.../yt-dlp.exe`   | `/usr/local/bin/yt-dlp`            |
| `FFMPEG_PATH`  | Ruta a la **carpeta** que contiene ffmpeg y ffprobe | `C:/.../WinGet/Links` | `/usr/bin`                         |
| `DOWNLOAD_DIR` | Carpeta de descargas                                | `downloads`           | `/var/www/downloader/downloads`    |
| `LOG_FILE`     | Archivo de logs                                     | `logs/app.log`        | `/var/www/downloader/logs/app.log` |

> 💡 **Tip:** `FFMPEG_PATH` apunta a la **carpeta** (no al .exe), porque yt-dlp busca `ffmpeg` **y** `ffprobe` dentro de ella.

---

## 🔧 Configuración

Copia el archivo de ejemplo:

```bash
cp .env.example .env
```

**`.env.example`:**

```env
# Entorno: development | production
APP_ENV=development

# Puerto del servidor
APP_PORT=8080

# Nivel de log: debug | info | warn | error
LOG_LEVEL=info

# Archivo de logs
LOG_FILE=logs/app.log

# Carpeta de descargas
DOWNLOAD_DIR=downloads

# Ruta al binario yt-dlp (o solo "yt-dlp" si está en el PATH)
YTDLP_PATH=yt-dlp

# Carpeta que contiene ffmpeg y ffprobe (vacío = usar PATH)
FFMPEG_PATH=
```

> ⚠️ **Go NO lee el `.env` automáticamente.** Exporta las variables antes de ejecutar, o integra `godotenv`. Ejemplos por SO:

**Windows (PowerShell):**

```powershell
$env:APP_PORT="8080"
$env:DOWNLOAD_DIR="downloads"
$env:YTDLP_PATH="yt-dlp"
$env:FFMPEG_PATH="C:\Users\<user>\AppData\Local\Microsoft\WinGet\Links"
go run ./cmd/api
```

**Linux / macOS (bash/zsh):**

```bash
export APP_PORT=8080
export DOWNLOAD_DIR=downloads
export YTDLP_PATH=/usr/local/bin/yt-dlp
export FFMPEG_PATH=/usr/bin
go run ./cmd/api
```

---

## ▶️ Ejecución

**Desarrollo:**

```bash
go mod tidy
go run ./cmd/api
```

**Compilar binario:**

```bash
go build -o bin/api ./cmd/api
./bin/api
```

**Windows:**

```powershell
go build -o bin/api.exe ./cmd/api
.\bin\api.exe
```

Deberías ver:

```json
{ "level": "INFO", "msg": "servidor escuchando", "addr": ":8080" }
```

---

## 🌐 Endpoints de la API

### `POST /api/v1/download/videos`

Descarga un video en la resolución indicada.

**Body:**

```json
{
  "url": "https://www.youtube.com/watch?v=XXXX",
  "format": "720p"
}
```

**Formatos permitidos:** `720p`, `1080p`

---

### `POST /api/v1/download/music`

Descarga el audio en el formato indicado.

**Body:**

```json
{
  "url": "https://www.youtube.com/watch?v=XXXX",
  "format": "mp3"
}
```

**Formatos permitidos:** `mp3`, `mp4`

---

### `GET /health`

Comprueba que el servicio está activo.

**Respuesta:**

```json
{ "status": "ok" }
```

---

### Respuestas

**Éxito:**

```json
{
  "success": true,
  "message": "Descarga completada",
  "data": {
    "file_name": "Titulo del video.mp4",
    "file_path": "/ruta/completa/Titulo del video.mp4",
    "format": "720p",
    "type": "video"
  }
}
```

**Error:**

```json
{
  "success": false,
  "error": "Formato incorrecto"
}
```

**Códigos de estado:**

| Código | Significado                                                           |
| ------ | --------------------------------------------------------------------- |
| `200`  | Descarga completada                                                   |
| `400`  | Error de validación (formato incorrecto, URL faltante, JSON inválido) |
| `405`  | Método no permitido                                                   |
| `500`  | Error interno (fallo de yt-dlp, ffmpeg, etc.)                         |

---

## 🧪 Ejemplos con curl

> **Video de prueba usado:** `https://www.youtube.com/watch?v=lSTQTc8EGs4`

### 🎬 Descargar video en 720p

**Linux / macOS:**

```bash
curl -X POST http://localhost:8080/api/v1/download/videos \
  -H "Content-Type: application/json" \
  -d '{"url":"https://www.youtube.com/watch?v=lSTQTc8EGs4","format":"720p"}'
```

**Windows (PowerShell — usa `curl.exe`):**

```powershell
curl.exe -X POST "http://localhost:8080/api/v1/download/videos" -H "Content-Type: application/json" -d "{\"url\":\"https://www.youtube.com/watch?v=lSTQTc8EGs4\",\"format\":\"720p\"}"
```

### 🎬 Descargar video en 1080p

```bash
curl -X POST http://localhost:8080/api/v1/download/videos \
  -H "Content-Type: application/json" \
  -d '{"url":"https://www.youtube.com/watch?v=lSTQTc8EGs4","format":"1080p"}'
```

```powershell
curl.exe -X POST "http://localhost:8080/api/v1/download/videos" -H "Content-Type: application/json" -d "{\"url\":\"https://www.youtube.com/watch?v=lSTQTc8EGs4\",\"format\":\"1080p\"}"
```

### 🎵 Descargar audio en MP3

```bash
curl -X POST http://localhost:8080/api/v1/download/music \
  -H "Content-Type: application/json" \
  -d '{"url":"https://www.youtube.com/watch?v=lSTQTc8EGs4","format":"mp3"}'
```

```powershell
curl.exe -X POST "http://localhost:8080/api/v1/download/music" -H "Content-Type: application/json" -d "{\"url\":\"https://www.youtube.com/watch?v=lSTQTc8EGs4\",\"format\":\"mp3\"}"
```

### 🎵 Descargar audio en MP4

```bash
curl -X POST http://localhost:8080/api/v1/download/music \
  -H "Content-Type: application/json" \
  -d '{"url":"https://www.youtube.com/watch?v=lSTQTc8EGs4","format":"mp4"}'
```

```powershell
curl.exe -X POST "http://localhost:8080/api/v1/download/music" -H "Content-Type: application/json" -d "{\"url\":\"https://www.youtube.com/watch?v=lSTQTc8EGs4\",\"format\":\"mp4\"}"
```

### 🩺 Health check

```bash
curl http://localhost:8080/health
```

### ⚠️ Ejemplo de formato inválido (devuelve error)

```bash
curl -X POST http://localhost:8080/api/v1/download/videos \
  -H "Content-Type: application/json" \
  -d '{"url":"https://www.youtube.com/watch?v=lSTQTc8EGs4","format":"480p"}'
```

**Respuesta:**

```json
{ "success": false, "error": "Formato incorrecto" }
```

---

## 🐧 Ejecución como servicio (systemd)

Solo para servidores Linux. Crea `/etc/systemd/system/downloader.service`:

```ini
[Unit]
Description=Video & Music Downloader API
After=network.target

[Service]
Type=simple
User=downloader
Group=downloader
WorkingDirectory=/var/www/downloader
Environment=APP_ENV=production
Environment=APP_PORT=8080
Environment=LOG_LEVEL=info
Environment=LOG_FILE=/var/www/downloader/logs/app.log
Environment=DOWNLOAD_DIR=/var/www/downloader/downloads
Environment=YTDLP_PATH=/usr/local/bin/yt-dlp
Environment=FFMPEG_PATH=/usr/bin
ExecStart=/var/www/downloader/bin/api
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
```

Activar:

```bash
sudo systemctl daemon-reload
sudo systemctl enable downloader
sudo systemctl start downloader
sudo systemctl status downloader

# Ver logs en vivo
sudo journalctl -u downloader -f
```

---

## 🔄 Mantenimiento

### Actualizar yt-dlp (IMPORTANTE)

YouTube cambia con frecuencia; yt-dlp se actualiza para adaptarse.

```bash
# Windows
winget upgrade yt-dlp

# Linux / macOS
yt-dlp -U
# o: brew upgrade yt-dlp
# o: pip install -U yt-dlp
```

**Automatizar en Linux (cron diario a las 4 AM):**

```bash
sudo crontab -e
```

```cron
0 4 * * * /usr/local/bin/yt-dlp -U >> /var/log/ytdlp-update.log 2>&1
```

### Actualizar ffmpeg

```bash
# Windows
winget upgrade Gyan.FFmpeg

# Linux
sudo apt update && sudo apt upgrade ffmpeg
```

---

## 🆘 Solución de problemas

### ❌ `yt-dlp: command not found`

El binario no está en el PATH. Usa la ruta absoluta en `.env`:

```env
YTDLP_PATH=/usr/local/bin/yt-dlp
```

### ❌ `ffprobe and ffmpeg not found`

`FFMPEG_PATH` está vacío o mal. Debe apuntar a la **carpeta**, no al `.exe`:

```env
FFMPEG_PATH=/usr/bin            # Linux/macOS
FFMPEG_PATH=C:/ffmpeg/bin       # Windows
```

### ❌ `No supported JavaScript runtime could be found`

Falta Node o Deno. Instálalo:

```bash
# Linux
sudo apt install -y nodejs
# Windows
winget install OpenJS.NodeJS
# macOS
brew install node
```

### ❌ El video se descarga SIN audio

Causas típicas:

1. **ffmpeg no encontrado** → revisa `FFMPEG_PATH`.
2. **Sin JS runtime** → instala Node/Deno.
3. Selección de formato incorrecta → el adaptador usa `bv*[height<=X]+ba/b[height<=X]`.

### ❌ `Formato incorrecto`

El valor de `format` no coincide **exactamente** (es case-sensitive):

- Video: solo `720p` o `1080p` (minúsculas).
- Música: solo `mp3` o `mp4` (minúsculas).

### ❌ La ruta devuelta viene duplicada

yt-dlp imprime varias líneas en `--print`. El adaptador debe tomar **solo la última línea no vacía**.

### ❌ `bind: address already in use`

El puerto está ocupado:

```bash
# Linux/macOS
lsof -i :8080
kill -9 <PID>

# Windows
netstat -ano | findstr :8080
taskkill /PID <PID> /F
```

### ✅ Comprobar que un archivo tiene audio

```bash
ffprobe -v error -select_streams a -show_entries stream=codec_type -of default=nw=1 archivo.mp4
```

Si sale `codec_type=audio` → tiene audio.

---

## 📄 Licencia

MIT
