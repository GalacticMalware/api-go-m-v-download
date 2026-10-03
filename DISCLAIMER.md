# ⚖️ Descargo de responsabilidad legal (Legal Disclaimer)

**Última actualización:** <fecha>

## 1. Naturaleza del software

Este proyecto (`Video & Music Downloader API`, en adelante "el Software") es una
**herramienta técnica genérica** de código abierto. El Software **no contiene, aloja,
distribuye ni incorpora** ningún contenido multimedia, ni incluye los binarios de
`yt-dlp`, `ffmpeg` o cualquier otro. El Software únicamente **invoca** herramientas
externas que el propio usuario instala en su sistema, bajo su exclusiva responsabilidad.

El Software **no posee ninguna afiliación, patrocinio o vínculo** con YouTube, Google
LLC, ni con ninguna otra plataforma, servicio o entidad mencionada.

## 2. Uso previsto

El Software está destinado **exclusivamente** a:

- Descargar contenido **libre de derechos de autor** (de dominio público o con
  licencia permisiva).
- Descargar contenido **sobre el que el usuario posea autorización explícita** de
  descarga por parte del titular de los derechos.
- Uso **personal, privado, educativo o de investigación** permitido por las leyes
  aplicables en la jurisdicción del usuario.

**Queda terminantemente prohibido** el uso del Software para:

- Descargar o distribuir material protegido por derechos de autor sin autorización.
- Redistribuir, revender, publicar o explotar comercialmente contenido descargado.
- Eludir o vulnerar medidas tecnológicas de protección cuando la ley lo prohíba.
- Cualquier actividad que infrinja los Términos de Servicio de plataformas terceras.
- Cualquier actividad ilegal conforme a la legislación aplicable.

## 3. Responsabilidad del usuario

**El usuario es el ÚNICO y exclusivo responsable** de:

- El uso que haga del Software.
- El contenido que descargue y su legalidad.
- El cumplimiento de los **Términos de Servicio** de YouTube y de cualquier otra
  plataforma.
- El cumplimiento de las **leyes de derechos de autor** de su país.
- El cumplimiento de las leyes **anti-elusión** (por ejemplo, DMCA §1201 en EE.UU.,
  Directiva 2001/29/CE en la UE u otras equivalentes).
- Cualquier daño o consecuencia legal derivada de su uso.

## 4. Exención de responsabilidad

Los autores, contribuyentes y titulares del copyright del Software:

- Se **eximen de toda responsabilidad** civil, penal o administrativa por el uso
  directo o indirecto del Software.
- **No garantizan** el funcionamiento correcto, ininterrumpido o libre de errores
  del Software.
- **No garantizan** que el Software sea apto para ningún propósito particular.
- **No asumen responsabilidad** por daños, pérdidas de datos, afectaciones a la
  reputación, sanciones o reclamaciones de terceros derivadas del uso del Software.

El Software se proporciona **"TAL CUAL" ("AS IS")**, **"SEGÚN DISPONIBILIDAD"**,
sin garantía de ningún tipo, expresa o implícita.

## 5. Naturaleza técnica y zona gris legal

El usuario reconoce y acepta que:

- La descarga de contenido de plataformas como YouTube **puede contravenir los
  Términos de Servicio** de dichas plataformas.
- Algunas jurisdicciones **prohíben eludir medidas técnicas de protección**, y las
  herramientas que descifran streams pueden caer en esa categoría.
- **Los autores no proporcionan asesoría legal** ni garantizan que el uso sea legal
  en todas las jurisdicciones.
- Las leyes varían según el país; el usuario debe **informarse y cumplir** con la
  legislación de su jurisdicción.

## 6. Dependencias de terceros

El Software **no redistribuye** los binarios de terceros. Al instalar y usar
`yt-dlp`, `ffmpeg`, `Go`, `Node.js` o `Deno`, el usuario queda sujeto a **sus
respectivas licencias**:

- **yt-dlp** → Licencia Unlicense (dominio público).
- **ffmpeg** → Licencia LGPL/GPL (según compilación).
- **Go** → Licencia BSD-3-Clause.

Los autores del Software **no son responsables** del cumplimiento de dichas
licencias por parte del usuario, ni de sus términos.

## 7. Indemnización

El usuario acepta **indemnizar y mantener indemnes** a los autores y contribuyentes
del Software frente a cualquier reclamación, demanda, daño, coste o gasto
(incluidos honorarios legales) que surja de su uso del Software o del incumplimiento
de estos términos.

## 8. Ausencia de garantía y limitación de responsabilidad

EN LA MÁXIMA MEDIDA PERMITIDA POR LA LEY APLICABLE, LOS AUTORES NO SERÁN
RESPONSABLES DE NINGÚN DAÑO DIRECTO, INDIRECTO, INCIDENTAL, ESPECIAL, EJEMPLAR
O CONSECUENTE (INCLUYENDO, SIN LIMITACIÓN, PÉRDIDA DE DATOS O BENEFICIOS, O
INTERRUPCIÓN DEL SERVICIO) DERIVADO DEL USO O LA IMPOSIBILIDAD DE USO DEL
SOFTWARE, CON INDEPENDENCIA DE SU FUNDAMENTO.

## 9. Modificaciones

Los autores se reservan el derecho de modificar este descargo en cualquier
momento. El uso continuado del Software tras dichas modificaciones constituye
su aceptación.

## 10. Aceptación

El uso, copia, instalación o distribución del Software **implica la aceptación
plena e incondicional** de este descargo de responsabilidad.

---

Este descargo **no es asesoría legal**. Si necesitas orientación jurídica
específica, consulta a un abogado de propiedad intelectual en tu jurisdicción.

````

---

## 3. Añadidos al `README.md`

Agrega estas secciones a tu README. Primero las **badges** al inicio:

```markdown README.md
# 🎬🎵 Video & Music Downloader API

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-blue)
![Status](https://img.shields.io/badge/Status-Educational%20Use-orange)

API REST escrita 100% en Go nativo (sin frameworks) para descargar videos de YouTube
en 720p/1080p y extraer audio en MP3/MP4.

> ⚠️ **AVISO LEGAL:** Este software es una herramienta **educativa/personal**.
> **No promueve la descarga de contenido con copyright.** El usuario es el único
> responsable de su uso. **Lee el [DISCLAIMER](./DISCLAIMER.md) completo antes de usarlo.**
````

Y **al final** del README:

```markdown README.md
---

## ⚖️ Aviso legal, licencia y uso

### 🚨 Uso previsto

Este proyecto se distribuye **exclusivamente con fines educativos, personales y de
investigación**. Está pensado para descargar contenido:

- De **dominio público** o libre de derechos.
- Sobre el que **el usuario tenga permiso explícito**.

### ❌ Usos prohibidos

- Descargar o redistribuir material **protegido por derechos de autor** sin autorización.
- Uso **comercial** sin licencias adecuadas.
- Cualquier uso que infrinja los **Términos de Servicio de YouTube** o las leyes
  de tu país.

### 📜 Descargo de responsabilidad

EL SOFTWARE SE PROPORCIONA **"TAL CUAL" ("AS IS")**, SIN GARANTÍA DE NINGÚN TIPO.
LOS AUTORES **NO SE HACEN RESPONSABLES** DEL USO QUE SE LE DÉ NI DE NINGÚN DAÑO
DERIVADO.

**El usuario es el ÚNICO responsable** de:
- Cumplir las leyes de derechos de autor de su jurisdicción.
- Cumplir los Términos de Servicio de las plataformas de terceros.
- Cualquier consecuencia legal, sanción o reclamación derivada de su uso.

📖 **Descargo completo y vinculante:** consulta el archivo **[DISCLAIMER.md](./DISCLAIMER.md)**.

### 📄 Licencia

Este proyecto está bajo la licencia **MIT**. Ver [LICENSE](./LICENSE).

### 🔗 Dependencias de terceros

Este proyecto **NO redistribuye** binarios externos. El usuario instala por su
cuenta `yt-dlp`, `ffmpeg`, `Go`, `Node.js`/`Deno`, sujetos a sus propias licencias.

| Herramienta | Licencia | Distribuida por este repo |
|-------------|----------|--------------------------|
| Este código | MIT | ✅ Sí |
| yt-dlp | Unlicense | ❌ No |
| ffmpeg | LGPL/GPL | ❌ No |
| Go | BSD-3-Clause | ❌ No |

---

_Los nombres de producto, marcas y plataformas mencionados (YouTube, Google, etc.)
son propiedad de sus respectivos titulares y se usan únicamente con fines
descriptivos e informativos. Este proyecto no está afiliado ni respaldado por ellos._
```

---

## 4. `TERMS.md` — Términos de uso (opcional pero recomendado)

```markdown TERMS.md
# Términos de Uso (Terms of Use)

Al descargar, instalar, ejecutar o usar este software ("el Software"),
aceptas los siguientes términos:

1. **Solo uso legal.** Te comprometes a usar el Software únicamente para fines
   lícitos y conforme a las leyes aplicables de tu jurisdicción.

2. **Responsabilidad exclusiva del usuario.** Eres el único responsable del uso
   del Software y del contenido que descargues.

3. **Respeto a derechos de autor.** No usarás el Software para descargar o
   redistribuir contenido protegido sin autorización de su titular.

4. **Términos de terceros.** Aceptas cumplir los Términos de Servicio de
   cualquier plataforma con la que el Software interactúe.

5. **Sin garantías.** El Software se entrega "TAL CUAL", sin garantías.

6. **Sin responsabilidad de los autores.** Los autores quedan exentos de toda
   responsabilidad por el uso del Software.

7. **Indemnización.** Indemnizarás a los autores frente a reclamaciones derivadas
   de tu uso indebido.

Si no aceptas estos términos, **no uses el Software.**
```

---

## 5. Estructura final del repo

```
video-music-downloader/
├── LICENSE              ← NUEVO (licencia MIT + advertencia extra)
├── DISCLAIMER.md        ← NUEVO (descargo legal completo)
├── TERMS.md             ← NUEVO (términos de uso)
├── README.md            ← ACTUALIZADO (badges + aviso legal + licencia)
├── .gitignore
├── go.mod
├── cmd/
├── internal/
├── pkg/
├── downloads/.gitkeep
└── logs/.gitkeep
```

---

## 6. `.gitignore` reforzado

Asegúrate de **no subir nunca** contenido descargado ni binarios de terceros:

```gitignore .gitignore
# ===== Binarios de la app =====
*.exe
/bin/
video-music-downloader
api
api.exe

# ===== Descargas (NUNCA subir contenido multimedia) =====
/downloads/*
!/downloads/.gitkeep

# ===== Logs =====
/logs/*
!/logs/.gitkeep

# ===== Variables de entorno (nunca subir secretos) =====
.env

# ===== Binarios externos (NO redistribuir) =====
yt-dlp
yt-dlp.exe
ffmpeg
ffmpeg.exe
ffprobe
ffprobe.exe

# ===== IDE / SO =====
.idea/
.vscode/
.DS_Store
Thumbs.db

# ===== Go =====
*.test
*.out
/vendor/
```
