const {createServer} = require("http");
const {Server} = require("socket.io");
const {Messages} = require("./db");
const jwt = require("jsonwebtoken");
const {pool} = require("./db");

const httpServer = createServer();
const io = new Server(httpServer, {
    cors: {
        origin: "*",
    },
});

io.on("connection", (socket) => {
    socket.on("SEND_MESSAGE_PRIVATE", async (data) => {
        const {token, chatId, message} = JSON.parse(data);
        try {
            const decoded = jwt.verify(token, "secret");


            // const result = await Messages().insert({
            //     chat_id: chatId,
            //     content: message,
            //     user_id: decoded.ID,
            // }).returning("*");
            // console.log({result})


            const client = await pool.connect();
            try {
                const result = await client.query(`INSERT INTO messages (private_chat_id, content, user_id)
                                                   VALUES ($1, $2, $3) RETURNING *`, [chatId, message, decoded.ID]);
                const {id, created_at, updated_at, deleted_at, user_id, content, private_chat_id} = result.rows[0]
                const normalizedMessage = {
                    ID: +id,
                    CreatedAt: created_at,
                    UpdatedAt: updated_at,
                    DeletedAt: deleted_at,
                    UserID: +user_id,
                    Content: content,
                    PrivateChatID: private_chat_id
                }
                    io.sockets.emit(`SEND_MESSAGE_PRIVATE:${chatId}`, JSON.stringify(normalizedMessage));
            } finally {
                client.release();
            }
        } catch (err) {
            console.error(err)
        }
    });
});

httpServer.listen(8082);

