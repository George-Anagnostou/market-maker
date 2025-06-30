# market-maker

`market-maker` is a turn-based CLI market making game. You, the player, enter a bid and ask price on each turn, and your inventory and cash positions update to reflect market conditions and your executions. After a series of turns, you score is calculated based on your inventory and cash posiitons. If you run out of cash, you lose!

## Market Maker Game Development Roadmap
### Overview
A single-player, turn-based CLI game evolving into a multiplayer, real-time web app. Players set bid-ask spreads to profit from orders, aiming to maximize cash/inventory. Built with Go for CLI and backend, TypeScript for web frontend, and PostgreSQL for user accounts/game history.

### Architecture

- CLI Phase: Go-based CLI app with game logic and terminal interface
- Web Phase: Client-server app:
    - Frontend: TypeScript + React
    - Backend: Go
    - Database: PostgreSQL (users, game history)
- Multiplayer: Players compete in shared markets, with real-time orders via WebSocket

### Milestones
#### Milestone 1: Core CLI Game Mechanics (Go)

Goal: Build a CLI game where the player sets bid-ask spreads, processes random orders, and tracks cash/inventory over a number of rounds.

Tasks:
- Create Go packages: player (cash, inventory, actions), market (order generation, trades), order (buy/sell, price, quantity).
- Implement random orders.
- Build turn-based loop: display state, input bid/ask, process orders.
- Add win/loss: Survive rounds or go bankrupt (cash < 0).

Deliverables:
Go CLI app (main.go, packages: player, market, order).
README.md with setup/run instructions.

#### Milestone 2: Enhanced Gameplay and CLI Polish

Goal: Add order generation factors, improve CLI UX, and log game history.
Tasks:
- Enhance orders: Add market trend (e.g., prices drift ±5%) or volatility.
- Improve CLI: Use color package for output, add trade summaries.
- Log history: Save rounds (cash, inventory, trades) to JSON file.
- Add scoring: Cash + inventory * average price.
- Test edge cases (e.g., negative inventory, wide spreads).

Deliverables:
- Updated Go app with enhanced orders and scoring.
- JSON logger.
- Polished CLI output.

#### Milestone 3: Backend Prototype (Go)

Goal: Build a Go backend with game logic and REST API, preparing for web integration.
Tasks:
- Refactor CLI: Separate game logic (game package) from interface (cli package).
- Create REST API: endpoints for game state (GET /state), set spread (POST /spread), next round (POST /next).
- Add PostgreSQL: Store user accounts (username, password hash) and game history.
- Implement basic authentication (JWT).
- Test API with CLI client.

Deliverables:
- Go backend (server.go) with REST API and PostgreSQL.
- Updated CLI to use backend API (optional).
- README with API docs.

#### Milestone 4: Web Frontend Prototype (TypeScript)

Goal: Port the game to a web app with a TypeScript/React frontend, keeping single-player mechanics.
Tasks:
- Set up TypeScript project.
- Create UI: Bid/ask inputs, cash/inventory display, trade history, next round button.
- Connect to Go backend via REST API.
- Add user login/signup (JWT-based).
- Test web gameplay: Match CLI mechanics.

Deliverables:
- TypeScript/React frontend (App.tsx).
- Basic web UI.
- Deployable app (e.g., Vercel).

#### Milestone 5: Multiplayer and Real-Time Features

Goal: Add multiplayer and real-time orders to the web app.
Tasks:
- Extend backend: Support multiple players per game session.
- Add real-time orders: Stream orders every periodically, resolving them periodically or in real-time.
- Update UI: Show anonymized competitor spreads, real-time trade feed.
- Balance multiplayer: Prioritize best spreads for order allocation.
- Test with 2-4 players.

Deliverables:
- Updated Go backend with multiplayer and WebSocket.
- Enhanced React UI.
- Multiplayer docs.

#### Milestone 6: Web Polish and Visuals

Goal: Add visualizations and polish the web app for sharing.
Tasks:
- Add charts: Use Chart.js for price history, inventory trends.
- Create order book UI: Show pending orders.
- Improve UX: Animations, responsive design, help modals.
- Deploy publicly.
- Test cross-browser compatibility.

Deliverables:
- Polished web app with charts/order book.
- Public URL.
- Final README with setup/play instructions.
