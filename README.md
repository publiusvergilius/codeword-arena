# 🎮 Codeword Arena

**Codeword Arena** is a real-time, multiplayer word association battle game built using Go and raw TCP sockets. Players face off in 1v1 duels, exchanging related words under pressure. Quick thinking and creativity are the keys to victory!

---

## ✨ Features

- 🔁 Turn-based 1v1 gameplay
- ⌛ Timed word association rounds
- ✅ Simple custom TCP protocol
- 🧠 Built-in synonym and keyword validation
- 📟 CLI-based UI (with potential for Wails or WebSocket upgrade)

---

## 🧩 How It Works

1. Players connect to the server via a TCP client.
2. The server matches two players into a game session.
3. Player 1 sends a codeword.
4. Player 2 has **10 seconds** to respond with a related word.
5. If the response is valid, they earn a point. Then roles switch.
6. First player to reach **5 points** wins the game.

---

## 🧪 Protocol Example

All messages are plain text over TCP using newline (`\n`) as the delimiter:

MATCHED ROUND_START:apple GUESS:fruit RESULT:correct SCORE:2-1 ROUND_START:dog GUESS:car RESULT:wrong SCORE:2-1 GAME_OVER:Player1 wins!
---

## 🚀 Getting Started

### 🔧 Requirements

- Go 1.21+
- Terminal or TCP client (like `telnet`, `netcat`, or a custom CLI)

### 🔨 Build & Run

#### 1. Clone the repo

```bash
git clone https://github.com/yourusername/codeword-arena.git
cd codeword-arena
