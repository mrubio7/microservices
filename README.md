# Backend - IBERCS

![Go Version](https://img.shields.io/badge/Go-1.23-blue) ![Build](https://img.shields.io/badge/Build-Passing-brightgreen)

Este es el repositorio del backend de **IBERCS**, diseñado para ofrecer una API RESTful escalable y modular. El proyecto está implementado en **Golang** utilizando arquitectura basada en **microservicios**.

## 🚀 Características principales

- **Microservicios**: Gestión separada de `Matches`, `Players`, `Teams`, `Tournaments` y `Users`.
- **Extensibilidad**: Fácil integración de nuevas funcionalidades como sistema de noticias y comentarios.
- **Autenticación**: OAuth2 para proteger los endpoints más críticos.
- **Modularidad**: Uso del patrón de repositorio genérico para consultas de base de datos.

## 🛠️ Tecnologías utilizadas

- **Lenguaje**: Golang
- **Framework HTTP**: Gin
- **Bases de datos**: PostgreSQL
- **Mensajería**: gRPC
- **Autenticación**: OAuth2 (Faceit)

## 🚦 Endpoints principales

### Usuarios (`/api/v2/user`)
- `GET /api/v2/user`: Recupera información del usuario actual.
- `PUT /api/v2/user`: Actualiza el perfil del usuario.
- `GET /api/v2/users/streams`: Recupera las transmisiones de un usuario.
- `POST /api/v2/auth/callback/faceit`: Endpoint de autenticación con Faceit.
- `POST /api/v2/auth`: Inicia sesión en la plataforma.
- `DELETE /api/v2/auth`: Cierra sesión en la plataforma.

### Jugadores (`/api/v2/player`)
- `GET /api/v2/player`: Obtiene información de un jugador.
- `GET /api/v2/players`: Lista todos los jugadores.
- `GET /api/v2/players/looking-for-team`: Lista los jugadores que buscan equipo.
- `POST /api/v2/players/looking-for-team`: Agrega un jugador como "buscando equipo".
- `PUT /api/v2/players/looking-for-team`: Actualiza la información de un jugador que busca equipo.
- `DELETE /api/v2/players/looking-for-team`: Elimina a un jugador de la lista de "buscando equipo".
- `GET /api/v2/players/prominent`: Obtiene los jugadores destacados.

### Equipos (`/api/v2/team`)
- `POST /api/v2/team/faceit`: Crea un equipo utilizando datos de Faceit.
- `GET /api/v2/team/faceit`: Obtiene un equipo desde Faceit.
- `GET /api/v2/team`: Recupera información sobre un equipo.
- `GET /api/v2/teams`: Lista todos los equipos.
- `GET /api/v2/teams/ranks`: Recupera el ranking de equipos.
- `GET /api/v2/teams/active`: Obtiene equipos activos.
- `GET /api/v2/team/player`: Recupera el equipo de un jugador.

### Torneos (`/api/v2/tournaments`)
- `GET /api/v2/tournaments`: Lista todos los torneos disponibles.
- `POST /api/v2/organizer`: Crea un nuevo organizador de torneos.
- `GET /api/v2/esea`: Obtiene información sobre las ligas ESEA.

### Partidos (`/api/v2/match`)
- `GET /api/v2/match`: Recupera información de un partido.
- `GET /api/v2/matches`: Lista todos los partidos.
- `GET /api/v2/matches/team`: Obtiene los partidos de un equipo específico.
- `POST /api/v2/match/stream`: Agrega un stream a un partido.
- `GET /api/v2/matches/range`: Recupera partidos dentro de un rango de fechas.


### Variables de entorno
Asegúrate de definir las siguientes variables de entorno en tu archivo `.env` o configuración del entorno de despliegue:
 ```bash
FACEIT_API_TOKEN=<tu-token> 
DB_HOST=<host-bd> 
DB_PORT=<puerto-bd> 
DB_USER=<usuario-bd> 
DB_PASSWORD=<contraseña-bd> 
DB_NAME=<nombre-bd>
```
(Preguntar por MD)

## 🔧 Instalación y ejecución

### Requisitos previos
- **Go** (v1.23+)
- **Docker Desktop** (Testing)

### Pasos
1. Clona el repositorio:
```bash
   git clone https://github.com/mrubio7/ibercs_backend.git
   cd ibercs_backend
```
2. Ejecuta los servicios
```bash
  make run # API gateway
  make ms-users # Microservicio Users
  make ms-players  # Microservicio Players
  make ms-teams  # Microservicio Teams
  make ms-tournaments # Microservicio Tournaments
  make ms-matches  # Microservicio Matches
```

